package shop

import (
	"dddemo/models"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type DomainShop struct {
	Name string
}

func NewShop() *DomainShop {
	return &DomainShop{
		Name: "Shop",
	}
}

func (d DomainShop) DomainName() string {
	return d.Name
}

func (d DomainShop) RootFolderPath() string {

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

	return fmt.Sprintf("./%s", rel)
}

func (d DomainShop) RootTemplatesFolder() string {
	return fmt.Sprintf("%s/web/templates", d.RootFolderPath())
}

func (d DomainShop) PathsToCSSFiles() []string {
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
func (d DomainShop) PathsToJSFiles() []string {
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

func (d DomainShop) SidebarMarkup() models.SidebarMarkup {
	return models.SidebarMarkup{
		Name: d.DomainName(),
		LI: []models.SidebarMarkupLI{
			{
				Name: "Hello",
				Href: "/tavern/shop/hello",
			},
			{
				Name: "Show Dishes",
				Href: "/tavern/dish/show",
			},
			{
				Name: "Create Dish",
				Href: "/tavern/dish/tmp_create",
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
