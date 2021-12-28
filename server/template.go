package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"dddemo/pkg/slice"
)

// parseTemplates parses .html files over base template in given directory and returns them.
func (serv *Server) parseTemplates(dir string) (map[string]*template.Template, error) {

	baseT := template.New("index.html").Funcs(map[string]interface{}{
		"hello":         func() string { return "HELLLLLOOOOL!" },
		"add":           func(x, y int) int { return x + y },
		"hasPermission": func(p string) bool { return true },
	})

	list, err := templatesFilesList(dir)
	if err != nil {
		return nil, err
	}

	// Добавляем частные темплейты из других директорий.
	baseT, err = baseT.ParseFiles(list...)
	if err != nil {
		return nil, fmt.Errorf("parsing templates: %v", err)
	}

	// template.C

	templates := make(map[string]*template.Template)
	fillTmps := func(tmpDir string) error {
		files, err := ioutil.ReadDir(tmpDir)
		if err != nil {
			return fmt.Errorf("read dir: %v", err)
		}

		for _, info := range files {
			if info.IsDir() || !slice.ContainsString(filepath.Ext(info.Name()), []string{".gohtml", ".html"}) || info.Name() == "index.gohtml" {
				continue
			}

			t, err := template.Must(baseT.Clone()).ParseFiles(filepath.Join(tmpDir, info.Name()))
			if err != nil {
				return fmt.Errorf("parse template %v: %v", info.Name(), err)
			}
			name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			if name == "" {
				return fmt.Errorf("template name cannot be empty")
			}

			if _, exists := templates[name]; exists {
				return fmt.Errorf("template %v is already added", name)
			}

			templates[name] = t
		}
		return nil
	}

	err = fillTmps(dir)
	if err != nil {
		return nil, fmt.Errorf("fillTmps: %v", err)
	}

	for _, d := range serv.domains {
		err = fillTmps(d.RootTemplatesFolder())
		if err != nil {
			return nil, fmt.Errorf("fillTmps: %v", err)
		}
	}

	return templates, nil
}

var ErrNoPermission = fmt.Errorf("you have no permissions for that")

func templatesFilesList(dir string) ([]string, error) {
	result := make([]string, 0, 185)
	p := filepath.Join(dir)
	err := filepath.Walk(p,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				result = append(result, path)
			}

			return nil
		})

	return result, err
}
