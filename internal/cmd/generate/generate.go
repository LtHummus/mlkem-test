package generate

import (
	"context"
	"crypto/mlkem"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var (
	Filename string
	Force    bool
)

func GenerateKey(ctx context.Context, cmd *cli.Command) error {
	dkFilename := fmt.Sprintf("%s.dk", Filename)
	ekFilename := fmt.Sprintf("%s.ek", Filename)

	_, err := os.Stat(dkFilename)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err == nil && !Force {
		return fmt.Errorf("%s exists. Use --force to overwrite", dkFilename)
	}

	_, err = os.Stat(ekFilename)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err == nil && !Force {
		return fmt.Errorf("%s exists. Use --force to overwrite", ekFilename)
	}

	dk, err := mlkem.GenerateKey768()
	if err != nil {
		return err
	}

	err = os.WriteFile(dkFilename, dk.Bytes(), 0600)
	if err != nil {
		return err
	}

	log.Printf("wrote %s", dkFilename)

	err = os.WriteFile(ekFilename, dk.EncapsulationKey().Bytes(), 0600)
	if err != nil {
		return err
	}

	log.Printf("wrote %s", ekFilename)

	return nil
}
