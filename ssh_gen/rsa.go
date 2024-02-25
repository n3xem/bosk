package ssh_gen

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/n3xem/bosk/entity"

	"golang.org/x/crypto/ssh"
)

const RSA_PEM_HEADER = "RSA PRIVATE KEY"

type RsaGenerator struct {
	Bits int
}

func (r RsaGenerator) GenerateKeyPair() (entity.SshKeyPair, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, r.Bits)
	if err != nil {
		return entity.SshKeyPair{}, err
	}

	privKeyPEM := &pem.Block{Type: RSA_PEM_HEADER, Bytes: x509.MarshalPKCS1PrivateKey(privKey)}
	privKeyBuf := new(bytes.Buffer)
	if err := pem.Encode(privKeyBuf, privKeyPEM); err != nil {
		return entity.SshKeyPair{}, err
	}

	pub, err := ssh.NewPublicKey(&privKey.PublicKey)
	if err != nil {
		return entity.SshKeyPair{}, err
	}
	pubKeyStr := string(ssh.MarshalAuthorizedKey(pub))

	return entity.SshKeyPair{
		PrivateKey: privKeyBuf.String(),
		PublicKey:  pubKeyStr,
	}, nil
}
