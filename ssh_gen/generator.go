package ssh_gen

import "github.com/n3xem/bosk/entity"

type SshKeyGenerator interface {
	GenerateKeyPair() (entity.SshKeyPair, error)
}
