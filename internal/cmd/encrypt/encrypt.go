package encrypt

import (
	"context"
	"fmt"
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

	fs, err := os.Stat(filename)
	if err != nil {
		return err
	}

	if fs.Size() == 0 {
		return fmt.Errorf("input file is empty")
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = encryption.Encrypt(keyBytes, uint64(fs.Size()), f, os.Stdout)
	if err != nil {
		return err
	}

	return nil
}
