/*
Copyright © 2024 yukyan
*/
package cmd

import (
	"fmt"

	"github.com/n3xem/bosk/entity"
	"github.com/n3xem/bosk/output"
	"github.com/n3xem/bosk/ssh_gen"
	"github.com/spf13/cobra"
)

var algorithm, fileType, fileName string
var numKeys int

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate ssh key pairs.",
	Long:  `generate ssh key pairs.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if algorithm != "rsa" && algorithm != "ed25519" {
			return fmt.Errorf("Unsupported algorithm: %s", algorithm)
		}
		if fileType != "json" && fileType != "csv" {
			return fmt.Errorf("Unsupported file type: %s", fileType)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var generator ssh_gen.SshKeyGenerator
		switch algorithm {
		case "rsa":
			generator = ssh_gen.RsaGenerator{
				Bits: 4096,
			}
		case "ed25519":
			generator = ssh_gen.Ed25519Generator{}
		}

		var oh output.OutputHandler
		switch fileType {
		case "json":
			oh = &output.JsonOutputHandlerImpl{
				FileName: fileName,
			}
		case "csv":
			oh = &output.CsvOutputHandlerImpl{
				FileName: fileName,
			}
		}

		pairs := []entity.SshKeyPair{}
		for i := 0; i < numKeys; i++ {
			pair, err := generator.GenerateKeyPair()
			if err != nil {
				return err
			}
			pairs = append(pairs, pair)
		}
		if err := oh.Output(pairs); err != nil {
			fmt.Println("Error writing output")
			return err
		} else {
			fmt.Printf("Generated %d key pairs\n", numKeys)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "rsa", "Algorithm to use for generating ssh key pairs. Supported algorithms are rsa, ed25519")
	genCmd.Flags().StringVarP(&fileType, "filetype", "t", "csv", "File type to use for storing ssh key pairs. Supported file types are json, csv")
	genCmd.Flags().StringVarP(&fileName, "filename", "f", "output.csv", "File name to store ssh key pairs")
	genCmd.Flags().IntVarP(&numKeys, "numkeys", "n", 1, "Number of key pairs to generate")

	genCmd.MarkFlagRequired("algorithm")
	genCmd.MarkFlagRequired("filetype")
	genCmd.MarkFlagFilename("filename")
	genCmd.MarkFlagRequired("numkeys")
}
