// Command qwsg-release-sign-offline is transferred to and run only on the
// approved custodian workstation. It never belongs on a QWSG runtime host.
package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"os"
	"runtime"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

const maxSigningInput = 1 << 20

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "offline signer:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 3 || args[0] != "sign" {
		return fmt.Errorf("usage: qwsg-release-sign-offline sign SIGNING_INPUT SIGNATURE_OUTPUT")
	}
	input, err := readRegular(args[1], maxSigningInput)
	if err != nil {
		return err
	}
	fmt.Fprint(os.Stderr, "OpenSSH private-key file: ")
	var keyPath string
	if _, err = fmt.Fscanln(os.Stdin, &keyPath); err != nil {
		return fmt.Errorf("private-key path unavailable")
	}
	keyPEM, err := readRegular(keyPath, 64<<10)
	if err != nil {
		return fmt.Errorf("private key unavailable")
	}
	defer zero(keyPEM)
	fmt.Fprint(os.Stderr, "Private-key passphrase: ")
	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Fprintln(os.Stderr)
	if err != nil {
		return fmt.Errorf("passphrase unavailable")
	}
	defer zero(passphrase)
	signature, err := sign(input, keyPEM, passphrase)
	if err != nil {
		return err
	}
	defer zero(signature)
	encoded := append([]byte(base64.StdEncoding.EncodeToString(signature)), '\n')
	return writeExclusive(args[2], encoded)
}

func sign(input, keyPEM, passphrase []byte) ([]byte, error) {
	parsed, err := ssh.ParseRawPrivateKeyWithPassphrase(keyPEM, passphrase)
	if err != nil {
		return nil, fmt.Errorf("encrypted OpenSSH key rejected")
	}
	privateKey, ok := parsed.(*ed25519.PrivateKey)
	if !ok || len(*privateKey) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("key is not Ed25519")
	}
	defer zero(*privateKey)
	return ed25519.Sign(*privateKey, input), nil
}

func readRegular(name string, maximum int64) ([]byte, error) {
	info, err := os.Lstat(name)
	if err != nil || !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 || info.Size() <= 0 || info.Size() > maximum {
		return nil, fmt.Errorf("unsafe or oversized input")
	}
	return os.ReadFile(name)
}

func writeExclusive(name string, content []byte) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}
	if runtime.GOOS == "windows" {
		// Windows ACLs remain a required custodian precondition; mode 0600 is
		// additionally enforced where the platform supports POSIX modes.
	}
	if _, err = file.Write(content); err == nil {
		err = file.Sync()
	}
	closeErr := file.Close()
	if err != nil {
		return err
	}
	return closeErr
}

func zero(value []byte) {
	for index := range value {
		value[index] = 0
	}
}
