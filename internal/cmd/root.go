package cmd

import (
	"github.com/urfave/cli/v3"

	"mlkem-test/internal/cmd/decrypt"
	"mlkem-test/internal/cmd/encrypt"
)

var RootCmd = &cli.Command{
	Commands: []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate a new mlkem key",
			Action:  GenerateKey,
		},
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "encrypt a file",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name:      "file",
					UsageText: "file to encrypt",
				},
			},
			Action: encrypt.Encrypt,
		},
		{
			Name:    "Decrypt",
			Aliases: []string{"d"},
			Usage:   "decrypt a file",
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name:      "file",
					UsageText: "file to decrypt",
				},
			},
			Action: decrypt.Decrypt,
		},
	},
}
