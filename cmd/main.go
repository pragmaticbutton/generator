package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:      "Generate new project",
		Args:      true,
		ArgsUsage: "ne znam",
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() == 0 {
				fmt.Println("missing name of the project")
				os.Exit(1)
			}

			if err := generateProject(ctx.Args().First()); err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func generateProject(name string) error {

	// TODO: improve this and make it configurable
	projLoc := ".."
	tmplDir := "templates/go/simple"

	err := filepath.WalkDir(tmplDir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		p, err1 := filepath.Rel(tmplDir, path)
		if err1 != nil {
			return err1
		}

		fp := filepath.Join(projLoc, name, p)

		if d.IsDir() {
			if err := os.Mkdir(fp, 0744); err != nil {
				return err
			}
		} else {
			f, err := os.Open(path)
			if err != nil {
				return err
			}

			fd, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}

			if _, err := io.Copy(fd, f); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
