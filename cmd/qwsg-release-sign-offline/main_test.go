package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"testing"

	"golang.org/x/crypto/ssh"
)

func TestEncryptedOpenSSHEd25519Signing(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	passphrase := []byte("non-production-test-passphrase")
	block, err := ssh.MarshalPrivateKeyWithPassphrase(privateKey, "NON-PRODUCTION TEST KEY", passphrase)
	if err != nil {
		t.Fatal(err)
	}
	input := []byte("deterministic non-production signing input")
	encoded := pem.EncodeToMemory(block)
	signature, err := sign(input, encoded, passphrase)
	if err != nil || !ed25519.Verify(publicKey, input, signature) {
		t.Fatalf("signature verification failed: %v", err)
	}
	if _, err = sign(input, encoded, []byte("wrong-passphrase")); err == nil {
		t.Fatal("wrong passphrase accepted")
	}
}
