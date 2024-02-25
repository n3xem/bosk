package ssh_gen

import (
	"bytes"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"

	"github.com/n3xem/bosk/entity"

	"golang.org/x/crypto/ssh"
)

const ED25519_PEM_HEADER = "OPENSSH PRIVATE KEY"

type Ed25519Generator struct {
}

func (e Ed25519Generator) GenerateKeyPair() (entity.SshKeyPair, error) {
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return entity.SshKeyPair{}, err
	}

	b, err := x509.MarshalPKCS8PrivateKey(privKey)
	if err != nil {
		return entity.SshKeyPair{}, err
	}
	privKeyPEM := &pem.Block{Type: ED25519_PEM_HEADER, Bytes: b}
	privKeyBuf := new(bytes.Buffer)
	if err := pem.Encode(privKeyBuf, privKeyPEM); err != nil {
		return entity.SshKeyPair{}, err
	}

	sshPubKey, err := ssh.NewPublicKey(pubKey)
	if err != nil {
		return entity.SshKeyPair{}, err
	}

	pubKeyStr := string(ssh.MarshalAuthorizedKey(sshPubKey))
	return entity.SshKeyPair{
		PrivateKey: privKeyBuf.String(),
		PublicKey:  pubKeyStr,
	}, nil
}
