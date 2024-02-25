package output

import (
	"encoding/csv"
	"os"

	"github.com/n3xem/bosk/entity"
)

type CsvOutputHandlerImpl struct {
	FileName string
}

func (c *CsvOutputHandlerImpl) Output(pairs []entity.SshKeyPair) error {
	records := [][]string{}
	for _, pair := range pairs {
		records = append(records, []string{pair.PrivateKey, pair.PublicKey})
	}
	f, err := os.Create(c.FileName)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	return w.WriteAll(records)
}
