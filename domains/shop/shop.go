package shop

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type DomainShop struct {
	Name string
}

func NewShop() *DomainShop {
	return &DomainShop{
		Name: "shop",
	}
}

func (s DomainShop) DomainName() string {
	return s.Name
}

func (s DomainShop) RootFolderPath() string {

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

func (s DomainShop) RootTemplatesFolder() string {
	return fmt.Sprintf("./%s/web/templates", s.RootFolderPath())
}
