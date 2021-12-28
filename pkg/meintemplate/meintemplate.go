package meintemplate

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/labstack/echo/v4"
)

type Templates map[string]*template.Template

// Render renders template with specified name.
func (ts Templates) Render(ectx echo.Context, code int, tmpl string, data interface{}) error {
	t, ok := ts[tmpl]
	if !ok {
		return fmt.Errorf("no template: %v", tmpl)
	}

	t, err := t.Clone()
	if err != nil {
		return fmt.Errorf("clone template: %v", err)
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return err
	}

	ectx.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ectx.Response().Header().Set("Pragma", "no-cache")
	ectx.Response().Header().Set("Expires", "0")
	ectx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	ectx.Response().WriteHeader(code)
	_, err = ectx.Response().Write(buf.Bytes())

	return err
}
