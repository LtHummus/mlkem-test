package cmd

import (
	"context"
	"crypto/mlkem"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func GenerateKey(ctx context.Context, cmd *cli.Command) error {
	dk, err := mlkem.GenerateKey768()
	if err != nil {
		return err
	}

	err = os.WriteFile("mlkem768.dk", dk.Bytes(), 0600)
	if err != nil {
		return err
	}

	log.Printf("wrote mlkem768.dk")

	err = os.WriteFile("mlkem768.ek", dk.EncapsulationKey().Bytes(), 0600)
	if err != nil {
		return err
	}

	log.Printf("wrote mlkem768.ek")

	return nil
}
