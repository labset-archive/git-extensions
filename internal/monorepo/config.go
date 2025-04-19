package monorepo

import (
	"fmt"
	"os"
	"path"
)

type View struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Includes    []string `json:"includes"`
}

func ensureViewsFolder() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: \n%w", err)
	}

	viewsDir := path.Join(wd, ".monorepo", "views")
	if err = os.MkdirAll(viewsDir, os.ModePerm); err != nil { //nolint:gosec
		return "", fmt.Errorf("failed to create views directory: \n%w", err)
	}

	return viewsDir, nil
}

func ensureViewFileCreated(name string) (*os.File, error) {
	viewsDir, err := ensureViewsFolder()
	if err != nil {
		return nil, fmt.Errorf("failed to ensure views folder: \n%w", err)
	}

	viewFile := path.Join(viewsDir, name+".json")
	file, err := os.Create(viewFile) //nolint:gosec

	if err != nil {
		return nil, fmt.Errorf("failed to create view file %s: \n%w", viewFile, err)
	}

	return file, nil
}

func ensureViewFileExists(name string) (*os.File, error) {
	viewsDir, err := ensureViewsFolder()
	if err != nil {
		return nil, fmt.Errorf("failed to ensure views folder: \n%w", err)
	}

	viewFile := path.Join(viewsDir, name+".json")
	file, err := os.Open(viewFile) //nolint:gosec

	if err != nil {
		return nil, fmt.Errorf("failed to open view file %s: \n%w", viewFile, err)
	}

	return file, nil
}
