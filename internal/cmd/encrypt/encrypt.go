package encrypt

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"

	"mlkem-test/internal/encryption"
)

func Encrypt(ctx context.Context, cmd *cli.Command) error {
	// read ek
	keyBytes, err := os.ReadFile("mlkem768.ek")
	if err != nil {
		return err
	}

	filename := cmd.Arguments[0].Get().(string)

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = encryption.Encrypt(keyBytes, f, os.Stdout)
	if err != nil {
		return err
	}

	return nil
}
