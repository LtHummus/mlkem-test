package decrypt

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"

	"mlkem-test/internal/encryption"
)

func Decrypt(ctx context.Context, cmd *cli.Command) error {
	dkBytes, err := os.ReadFile("mlkem768.dk")
	if err != nil {
		return err
	}

	filename := cmd.Arguments[0].Get().(string)
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = encryption.Decrypt(dkBytes, f, os.Stdout)
	if err != nil {
		return err
	}

	return nil
}
