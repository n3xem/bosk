package output

import (
	"encoding/json"
	"os"

	"github.com/n3xem/bosk/entity"
)

type JsonOutputHandlerImpl struct {
	FileName string
}

func (j *JsonOutputHandlerImpl) Output(pairs []entity.SshKeyPair) error {
	jsonData, err := json.Marshal(pairs)
	if err != nil {
		return err
	}
	return os.WriteFile(j.FileName, jsonData, 0644)
}
