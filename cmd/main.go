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
		// TODO: improve these parameters
		Name:                 "Generate new project",
		Usage:                "go run cmd/main.go newProjectName",
		Args:                 true,
		ArgsUsage:            "Specify new project's name",
		EnableBashCompletion: true,
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
	location := ".."
	templateDir := "templates/go/simple"

	if err := copyContentsOfDir(templateDir, location, name); err != nil {
		return err
	}

	return nil
}

func copyContentsOfDir(dirPath, location, name string) error {
	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		relPath, relErr := filepath.Rel(dirPath, path)
		if relErr != nil {
			return relErr
		}

		destPath := filepath.Join(location, name, relPath)

		if d.IsDir() {
			if err := os.Mkdir(destPath, 0744); err != nil {
				return err
			}
		} else {
			src, err := os.Open(path)
			if err != nil {
				return err
			}

			dst, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}

			if _, err := io.Copy(dst, src); err != nil {
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
