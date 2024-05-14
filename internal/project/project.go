package project

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
)

const (
	templateDir = "templates/go/server"
)

// Project represents project which should be generated
type Project struct {
	Location string
	Name     string
}

func New(location, name string) *Project {
	return &Project{
		Location: location,
		Name:     name,
	}
}

func (p *Project) Generate() error {

	if err := p.validate(); err != nil {
		return err
	}

	err := filepath.WalkDir(templateDir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}

		relPath, relErr := filepath.Rel(templateDir, path)
		if relErr != nil {
			return relErr
		}

		destPath := filepath.Join(p.Location, p.Name, relPath)

		if d.IsDir() {
			if err := os.Mkdir(destPath, 0744); err != nil {
				return err
			}
		} else {

			dst, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return err
			}

			tmpl, err := template.ParseFiles(path)
			if err != nil {
				return err
			}

			err = tmpl.Execute(dst, p)
			if err != nil {
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

func (p *Project) validate() error {

	d, err := os.Stat(p.Location)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%v doesn't exist", p.Location)
		}
		return err
	}

	if !d.IsDir() {
		return fmt.Errorf("%s is not a valid directory path", p.Location)
	}

	_, err = os.Open(filepath.Join(p.Location, p.Name))
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return fmt.Errorf("project %s already exists in %s", p.Name, p.Location)
}
