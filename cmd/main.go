package main

import (
	"fmt"
	"generator/internal/project"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:                 "Generator",
		Usage:                "Generates new project out of existing templates",
		UsageText:            "go run ./... [location] [name]",
		Args:                 true,
		EnableBashCompletion: true,
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() != 2 {
				return fmt.Errorf("expected number of parameters is 2")
			}

			location := ctx.Args().Get(0)
			name := ctx.Args().Get(1)
			p := project.New(location, name)
			if err := p.Generate(); err != nil {
				return fmt.Errorf("%v", err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
