package output

import "github.com/n3xem/bosk/entity"

type OutputHandler interface {
	Output(pairs []entity.SshKeyPair) error
}
