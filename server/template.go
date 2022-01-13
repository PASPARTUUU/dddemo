package server

import (
	"embed"
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"dddemo/models"
	"dddemo/pkg/slice"
)

type Dmns struct {
	SidebarMarkup []models.SidebarMarkup
	CSSFiles      []string
	JSFiles       []string
}

func (srv *Server) domainsInit() Dmns {
	var dmns Dmns

	for _, d := range srv.domains {
		dmns.SidebarMarkup = append(dmns.SidebarMarkup, d.SidebarMarkup())
		dmns.CSSFiles = append(dmns.CSSFiles, d.PathsToCSSFiles()...)
		dmns.JSFiles = append(dmns.JSFiles, d.PathsToJSFiles()...)
	}

	return dmns
}

// parseTemplates parses .html files over base template in given directory and returns them.
func (srv *Server) parseTemplates(dir string) (map[string]*template.Template, error) {

	baseT := template.New("index.html").Funcs(map[string]interface{}{
		"hello":         func() string { return "HELLLLLOOOOL!" },
		"add":           func(x, y int) int { return x + y },
		"hasPermission": func(p string) bool { return true },
		"domainsInit":   srv.domainsInit,
	})

	// Добавляем частные темплейты из других директорий.
	baseT, err := baseT.ParseFS(srv.embedWeb, dir+"/*")
	if err != nil {
		return nil, err
	}

	templates := make(map[string]*template.Template)
	fillTmps := func(fs embed.FS, tmpDir string) error {

		filesInfo, err := fs.ReadDir(tmpDir)
		if err != nil {
			return err
		}

		for _, info := range filesInfo {
			if info.IsDir() || !slice.ContainsString(filepath.Ext(info.Name()), []string{".gohtml", ".html"}) || info.Name() == "index.gohtml" {
				continue
			}

			t, err := template.Must(baseT.Clone()).ParseFS(fs, filepath.Join(tmpDir, info.Name()))
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

	err = fillTmps(srv.embedWeb, dir)
	if err != nil {
		return nil, fmt.Errorf("fillTmps: %v", err)
	}

	for _, d := range srv.domains {
		err = fillTmps(srv.embedDomains, d.RootTemplatesFolder())
		if err != nil {
			return nil, fmt.Errorf("fillTmps: %v", err)
		}
	}

	return templates, nil
}

var ErrNoPermission = fmt.Errorf("you have no permissions for that")
