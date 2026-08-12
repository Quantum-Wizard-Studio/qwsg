package smtpnotification

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"net"
	"strings"
	"syscall"
	"testing"
	"time"

	"quantumwizard.hu/qwsg/internal/alert"
	"quantumwizard.hu/qwsg/internal/configuration"
	"quantumwizard.hu/qwsg/internal/notification"
)

func effective(fields map[string]string) configuration.Effective {
	return configuration.Effective{Values: configuration.Model{Extensions: []configuration.Extension{{ID: ExtensionID, Version: "1.0", Fields: fields}}}}
}

func TestControlledImplicitTLSSMTP(t *testing.T) {
	cert, pool := testCertificate(t)
	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12})
	if err != nil {
		if errors.Is(err, syscall.EPERM) {
			t.Skip("loopback listeners prohibited by execution sandbox")
		}
		t.Fatal(err)
	}
	defer listener.Close()
	done := make(chan error, 1)
	go func() {
		conn, e := listener.Accept()
		if e != nil {
			done <- e
			return
		}
		defer conn.Close()
		rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		write := func(s string) { _, _ = rw.WriteString(s + "\r\n"); _ = rw.Flush() }
		write("220 test ESMTP")
		data := false
		for {
			line, e := rw.ReadString('\n')
			if e != nil {
				done <- e
				return
			}
			line = strings.TrimSpace(line)
			if data {
				if line == "." {
					data = false
					write("250 accepted")
				}
				continue
			}
			upper := strings.ToUpper(line)
			switch {
			case strings.HasPrefix(upper, "EHLO"):
				write("250-test")
				write("250 SIZE 4096")
			case strings.HasPrefix(upper, "MAIL FROM:"):
				write("250 ok")
			case strings.HasPrefix(upper, "RCPT TO:"):
				write("250 ok")
			case upper == "DATA":
				data = true
				write("354 end data")
			case upper == "QUIT":
				write("221 bye")
				done <- nil
				return
			default:
				write("250 ok")
			}
		}
	}()
	port := listener.Addr().(*net.TCPAddr).Port
	cfg := Config{Enabled: true, Recipient: "admin@example.test", Host: "localhost", Port: port, Sender: "qwsg@example.test", Security: "implicit_tls", Auth: "none", Timeout: 5 * time.Second}
	now := time.Now().UTC()
	req := notification.Request{DestinationRef: "community.admin", Deadline: now.Add(5 * time.Second), Envelope: notification.DeliveryEnvelope{Event: alert.EventEntered, Severity: alert.Warning, Category: alert.EngineeringCondition, ReasonToken: "test", EventTime: now, EvidenceReferences: []string{}}}
	result := Provider{Config: cfg, TLSConfig: &tls.Config{RootCAs: pool, MinVersion: tls.VersionTLS12}}.Deliver(context.Background(), req)
	if result.Status != notification.StatusAccepted {
		t.Fatalf("result %+v", result)
	}
	if err := <-done; err != nil {
		t.Fatal(err)
	}
}

func testCertificate(t *testing.T) (tls.Certificate, *x509.CertPool) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	tmpl := x509.Certificate{SerialNumber: big.NewInt(1), Subject: pkix.Name{CommonName: "localhost"}, NotBefore: now.Add(-time.Hour), NotAfter: now.Add(time.Hour), KeyUsage: x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, DNSNames: []string{"localhost"}}
	der, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		t.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(certPEM)
	return cert, pool
}
func TestCommunityConfigurationAndCardinality(t *testing.T) {
	valid := map[string]string{"enabled": "true", "recipients": "admin@example.test", "host": "smtp.example.test", "port": "465", "sender": "qwsg@example.test", "security": "implicit_tls", "auth": "password", "username": "admin", "credential_ref": "smtp.admin", "timeout": "5s"}
	c, err := FromEffective(effective(valid))
	if err != nil || !c.Enabled || c.Recipient != "admin@example.test" {
		t.Fatalf("valid: %+v %v", c, err)
	}
	bad := map[string]string{}
	for k, v := range valid {
		bad[k] = v
	}
	bad["recipients"] = "a@example.test,b@example.test"
	if _, err = FromEffective(effective(bad)); err == nil {
		t.Fatal("multiple Community recipients accepted")
	}
}
func TestPreflightClassifications(t *testing.T) {
	f := Preflight(Config{Enabled: true, Auth: "password"}, false)
	if len(f) != 4 || f[1].Classification != MissingRequired || f[2].Classification != UnknownVerification || Ready(f) {
		t.Fatalf("findings %+v", f)
	}
	disabled := Preflight(Config{}, false)
	if disabled[0].Classification != MissingOptional || !Ready(disabled) {
		t.Fatalf("disabled %+v", disabled)
	}
}
func TestPolicyAndPrivacySafeRender(t *testing.T) {
	c := Config{Enabled: true, Recipient: "admin@example.test", CredentialRef: "smtp.admin"}
	p, err := Policy(c)
	if err != nil || len(p.Routes) != 1 || p.Retry.MaxAttempts != 3 {
		t.Fatalf("policy %+v %v", p, err)
	}
	now := time.Date(2026, 8, 12, 12, 0, 0, 0, time.UTC)
	req := notification.Request{DestinationRef: c.Recipient, Envelope: notification.DeliveryEnvelope{Event: alert.EventEntered, Severity: alert.Warning, Category: alert.EngineeringCondition, ReasonToken: "policy_attention", EventTime: now}}
	message := string(Render(req))
	for _, forbidden := range []string{"password", "192.168.", "/etc/", "smtp.admin"} {
		if strings.Contains(message, forbidden) {
			t.Fatalf("leaked %q", forbidden)
		}
	}
	for _, required := range []string{"QWSG warning: entered", "policy_attention", "2026-08-12T12:00:00Z"} {
		if !strings.Contains(message, required) {
			t.Fatalf("missing %q", required)
		}
	}
}
