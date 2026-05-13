package cmd

import (
	"github.com/urfave/cli/v3"

	"mlkem-test/internal/cmd/decrypt"
	"mlkem-test/internal/cmd/encrypt"
	"mlkem-test/internal/cmd/generate"
)

var RootCmd = &cli.Command{
	Commands: []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "generate a new mlkem key",
			Action:  generate.GenerateKey,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "key-name",
					Value:       "mlkem768",
					Usage:       "keys will be written as <name>.dk and <name>.ek",
					Destination: &generate.Filename,
				},
				&cli.BoolFlag{
					Name:        "force",
					Value:       false,
					Usage:       "overwrite keys if they exist",
					Destination: &generate.Force,
				},
			},
		},
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "encrypt a file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "key-file",
					Usage:       "encapsulation key to use",
					DefaultText: "mlkem768.ek",
					Value:       "mlkem768.ek",
					Destination: &encrypt.KeyFile,
				},
			},
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name:        "file",
					UsageText:   "file to encrypt",
					Destination: &encrypt.InputFile,
				},
			},
			Action: encrypt.Encrypt,
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "decrypt a file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "key-file",
					Usage:       "dencapsulation key to use",
					DefaultText: "mlkem768.dk",
					Value:       "mlkem768.dk",
					Destination: &decrypt.KeyFile,
				},
			},
			Arguments: []cli.Argument{
				&cli.StringArg{
					Name:        "file",
					UsageText:   "file to decrypt",
					Destination: &decrypt.InputFile,
				},
			},
			Action: decrypt.Decrypt,
		},
	},
}
