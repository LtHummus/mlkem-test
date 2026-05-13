package decrypt

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"

	"mlkem-test/internal/encryption"
)

var (
	InputFile string
	KeyFile   string
)

func Decrypt(ctx context.Context, cmd *cli.Command) error {
	dkBytes, err := os.ReadFile(KeyFile)
	if err != nil {
		return err
	}

	f, err := os.Open(InputFile)
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
