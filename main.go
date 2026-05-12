package main

import (
	"context"
	"log"
	"os"

	"mlkem-test/internal/cmd"
)

func main() {
	if err := cmd.RootCmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
