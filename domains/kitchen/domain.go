package kitchen

import (
	"dddemo/models"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type DomainKitchenMeta struct {
	Name string
}

func NewKitchen() *DomainKitchenMeta {
	return &DomainKitchenMeta{
		Name: "Kitchen",
	}
}

func (d DomainKitchenMeta) DomainName() string {
	return d.Name
}

func (d DomainKitchenMeta) RootFolderPath() string {

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("not ok")
	}

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	rel, err := filepath.Rel(dir, filepath.Dir(file))
	if err != nil {
		panic(err)
	}

	return rel
}

func (d DomainKitchenMeta) RootTemplatesFolder() string {
	return fmt.Sprintf("%s/web/templates", d.RootFolderPath())
}
func (d DomainKitchenMeta) RootStaticFolder() string {
	return fmt.Sprintf("%s/web/static", d.RootFolderPath())
}

func (d DomainKitchenMeta) PathsToCSSFiles() []string {
	var res []string

	files, err := filesList(fmt.Sprintf("./%s/web/static", d.RootFolderPath()))
	if err != nil {
		panic(err)
	}

	for _, v := range files {
		if path.Ext(v) == ".css" {
			res = append(res, fmt.Sprintf("/%s", v))
		}
	}

	return res
}
func (d DomainKitchenMeta) PathsToJSFiles() []string {
	var res []string

	files, err := filesList(fmt.Sprintf("./%s/web/static", d.RootFolderPath()))
	if err != nil {
		panic(err)
	}

	for _, v := range files {
		if path.Ext(v) == ".js" {
			res = append(res, fmt.Sprintf("/%s", v))
		}
	}

	return res
}

func (d DomainKitchenMeta) SidebarMarkup() models.SidebarMarkup {
	return models.SidebarMarkup{
		Name: d.DomainName(),
		LI: []models.SidebarMarkupLI{
			{
				Name: "Hello",
				Href: "/tavern/kitchen/hello",
			},
		},
	}
}

func filesList(dir string) ([]string, error) {
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
