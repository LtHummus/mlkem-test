package encrypt

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"mlkem-test/internal/encryption"
)

var (
	InputFile string
	KeyFile   string
)

func Encrypt(ctx context.Context, cmd *cli.Command) error {
	// read ek
	keyBytes, err := os.ReadFile(KeyFile)
	if err != nil {
		return err
	}

	fs, err := os.Stat(InputFile)
	if err != nil {
		return err
	}

	if fs.Size() == 0 {
		return fmt.Errorf("input file is empty")
	}

	f, err := os.Open(InputFile)
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
