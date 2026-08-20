// Package smtpnotification activates Community email behind the canonical Notification provider.
package smtpnotification

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/assessment"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/notification"
)

const ExtensionID = "notification.email"
const ProviderID = "community.smtp"

type Config struct {
	Enabled                                                          bool
	Recipient, Host, Sender, Security, Auth, Username, CredentialRef string
	Port                                                             int
	Timeout                                                          time.Duration
}
type Classification = assessment.Classification

const (
	Satisfied           = assessment.Satisfied
	MissingRequired     = assessment.MissingRequired
	MissingOptional     = assessment.MissingOptional
	UnknownVerification = assessment.UnknownVerification
	Incompatible        = assessment.Incompatible
)

// Finding retains the Task 046 focused JSON contract while its classification
// is now owned by the common assessment model.
type Finding struct {
	Requirement    string                    `json:"requirement"`
	Classification assessment.Classification `json:"classification"`
	Remediation    string                    `json:"remediation,omitempty"`
}

func FromEffective(e configuration.Effective) (Config, error) {
	c := Config{Security: "starttls", Auth: "none", Timeout: 10 * time.Second}
	var fields map[string]string
	for _, x := range e.Values.Extensions {
		if x.ID == ExtensionID {
			fields = x.Fields
			break
		}
	}
	if fields == nil {
		return c, nil
	}
	allowed := map[string]bool{"enabled": true, "recipients": true, "host": true, "port": true, "sender": true, "security": true, "auth": true, "username": true, "credential_ref": true, "timeout": true}
	for k := range fields {
		if !allowed[k] {
			return c, fmt.Errorf("unsupported email field")
		}
	}
	c.Enabled = fields["enabled"] == "true"
	if fields["enabled"] != "" && fields["enabled"] != "true" && fields["enabled"] != "false" {
		return c, fmt.Errorf("invalid email enabled value")
	}
	recipients := strings.Split(fields["recipients"], ",")
	if fields["recipients"] != "" && len(recipients) != 1 {
		return c, fmt.Errorf("Community email requires exactly one recipient")
	}
	if len(recipients) == 1 {
		c.Recipient = recipients[0]
	}
	c.Host, c.Sender, c.Username, c.CredentialRef = fields["host"], fields["sender"], fields["username"], fields["credential_ref"]
	if fields["security"] != "" {
		c.Security = fields["security"]
	}
	if fields["auth"] != "" {
		c.Auth = fields["auth"]
	}
	if fields["port"] != "" {
		n, err := strconv.Atoi(fields["port"])
		if err != nil {
			return c, fmt.Errorf("invalid SMTP port")
		}
		c.Port = n
	}
	if fields["timeout"] != "" {
		d, err := time.ParseDuration(fields["timeout"])
		if err != nil {
			return c, fmt.Errorf("invalid SMTP timeout")
		}
		c.Timeout = d
	}
	if !c.Enabled {
		return c, nil
	}
	if c.Host == "" || strings.ContainsAny(c.Host, "/\\ \t\r\n") || c.Port < 1 || c.Port > 65535 || c.Timeout < time.Second || c.Timeout > time.Minute {
		return c, fmt.Errorf("invalid SMTP endpoint")
	}
	if _, err := mail.ParseAddress(c.Recipient); err != nil {
		return c, fmt.Errorf("invalid Community recipient")
	}
	if _, err := mail.ParseAddress(c.Sender); err != nil {
		return c, fmt.Errorf("invalid SMTP sender")
	}
	if c.Security != "starttls" && c.Security != "implicit_tls" {
		return c, fmt.Errorf("TLS is required")
	}
	if c.Auth != "none" && c.Auth != "password" {
		return c, fmt.Errorf("unsupported SMTP authentication")
	}
	if c.Auth == "password" && (c.Username == "" || c.CredentialRef == "") {
		return c, fmt.Errorf("SMTP credential reference required")
	}
	return c, nil
}

func Preflight(c Config, credentialAvailable bool) []Finding {
	f := []Finding{}
	add := func(r string, x Classification) { f = append(f, Finding{Requirement: r, Classification: x}) }
	if !c.Enabled {
		add("notification_enabled", MissingOptional)
		return f
	}
	add("smtp_configuration", Satisfied)
	if c.Auth == "password" && !credentialAvailable {
		add("smtp_credential", MissingRequired)
	} else {
		add("smtp_credential", Satisfied)
	}
	add("tls_trust", UnknownVerification)
	add("smtp_network", UnknownVerification)
	return f
}
func Ready(findings []Finding) bool {
	for _, f := range findings {
		if f.Classification == MissingRequired || f.Classification == Incompatible {
			return false
		}
	}
	return true
}

type Provider struct {
	Config   Config
	Password []byte
	// TLSConfig is an injected trust boundary for controlled tests. Production
	// callers leave it nil and therefore use the system trust store.
	TLSConfig *tls.Config
}

func (p Provider) Descriptor() notification.ProviderDescriptor {
	return notification.ProviderDescriptor{SchemaName: notification.ProviderSchema, SchemaVersion: notification.SchemaVersion, ID: ProviderID, Channels: []notification.Channel{notification.Email}}
}
func (p Provider) Deliver(ctx context.Context, req notification.Request) notification.ProviderResult {
	now := time.Now().UTC()
	result := notification.ProviderResult{SchemaName: notification.ProviderResultSchema, SchemaVersion: notification.SchemaVersion, CompletedAt: now, EvidenceTokens: []string{"smtp_attempted"}}
	err := p.send(ctx, req)
	if err == nil {
		result.Status = notification.StatusAccepted
		result.Failure = notification.FailureNone
		result.EvidenceTokens = []string{"smtp_server_accepted"}
		return result
	}
	result.Status = notification.StatusRetryableFailure
	result.Failure = notification.FailureRetryable
	result.EvidenceTokens = []string{"smtp_delivery_failed"}
	var op *net.OpError
	if errors.As(err, &op) && !op.Timeout() {
		return result
	}
	if strings.Contains(err.Error(), "authentication") || strings.Contains(err.Error(), "certificate") {
		result.Status = notification.StatusTerminalFailure
		result.Failure = notification.FailureAuthentication
	}
	return result
}
func (p Provider) send(ctx context.Context, req notification.Request) error {
	deadline := time.Now().Add(p.Config.Timeout)
	if d, ok := ctx.Deadline(); ok && d.Before(deadline) {
		deadline = d
	}
	d := net.Dialer{Timeout: p.Config.Timeout}
	raw, err := d.DialContext(ctx, "tcp", net.JoinHostPort(p.Config.Host, strconv.Itoa(p.Config.Port)))
	if err != nil {
		return err
	}
	defer raw.Close()
	_ = raw.SetDeadline(deadline)
	tlsCfg := &tls.Config{ServerName: p.Config.Host, MinVersion: tls.VersionTLS12}
	if p.TLSConfig != nil {
		tlsCfg = p.TLSConfig.Clone()
		tlsCfg.ServerName = p.Config.Host
		if tlsCfg.MinVersion < tls.VersionTLS12 {
			tlsCfg.MinVersion = tls.VersionTLS12
		}
	}
	var conn net.Conn = raw
	if p.Config.Security == "implicit_tls" {
		tc := tls.Client(raw, tlsCfg)
		if err = tc.HandshakeContext(ctx); err != nil {
			return err
		}
		conn = tc
	}
	c, err := smtp.NewClient(conn, p.Config.Host)
	if err != nil {
		return err
	}
	defer c.Close()
	if p.Config.Security == "starttls" {
		if ok, _ := c.Extension("STARTTLS"); !ok {
			return fmt.Errorf("STARTTLS unavailable")
		}
		if err = c.StartTLS(tlsCfg); err != nil {
			return err
		}
	}
	if p.Config.Auth == "password" {
		if err = c.Auth(smtp.PlainAuth("", p.Config.Username, string(p.Password), p.Config.Host)); err != nil {
			return fmt.Errorf("authentication failed")
		}
	}
	if err = c.Mail(address(p.Config.Sender)); err != nil {
		return err
	}
	if err = c.Rcpt(address(p.Config.Recipient)); err != nil {
		return err
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(renderMessage(p.Config, req))
	if closeErr := w.Close(); err == nil {
		err = closeErr
	}
	if err == nil {
		err = c.Quit()
	}
	return err
}
func address(v string) string {
	a, err := mail.ParseAddress(v)
	if err != nil {
		return ""
	}
	return a.Address
}
func Render(req notification.Request) []byte {
	return renderMessage(Config{Sender: req.DestinationRef, Recipient: req.DestinationRef}, req)
}
func renderMessage(config Config, req notification.Request) []byte {
	e := req.Envelope
	subject := fmt.Sprintf("QWSG %s: %s", e.Severity, e.Event)
	body := fmt.Sprintf("QWSG detected an operator-attention event.\r\n\r\nSeverity: %s\r\nCategory: %s\r\nEvent: %s\r\nReason: %s\r\nTime: %s\r\n\r\nInspect detailed evidence locally with QWSG.\r\n", e.Severity, e.Category, e.Event, e.ReasonToken, e.EventTime.UTC().Format(time.RFC3339))
	return []byte("From: " + config.Sender + "\r\nTo: " + config.Recipient + "\r\nSubject: " + subject + "\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n" + body)
}
func Policy(c Config) (notification.Policy, error) {
	if !c.Enabled {
		return notification.NewPolicy(notification.RetryPolicy{MaxAttempts: 1, DeliveryWindowNS: int64(time.Hour), BackoffNS: []int64{}}, []notification.Route{}, []notification.EndpointReference{}, []notification.ProviderBinding{})
	}
	r := notification.Route{ID: "community.email", Enabled: true, Severities: []alert.Severity{alert.Warning, alert.Critical, alert.Emergency}, Categories: []alert.Category{}, Events: []alert.EventKind{alert.EventEntered, alert.EventEscalated, alert.EventRecovered}, EndpointIDs: []string{"community.admin"}}
	e := notification.EndpointReference{ID: "community.admin", Channel: notification.Email, DestinationRef: "community.admin", SecretRef: c.CredentialRef}
	b := notification.ProviderBinding{ID: "community.smtp", Channel: notification.Email, ProviderID: ProviderID}
	return notification.NewPolicy(notification.RetryPolicy{MaxAttempts: 3, DeliveryWindowNS: int64(time.Hour), BackoffNS: []int64{int64(time.Minute), int64(5 * time.Minute)}}, []notification.Route{r}, []notification.EndpointReference{e}, []notification.ProviderBinding{b})
}
