package kitchen

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type DomainKitchen struct {
	Name string
}

func NewKitchen() *DomainKitchen {
	return &DomainKitchen{
		Name: "kitchen",
	}
}

func (s DomainKitchen) DomainName() string {
	return s.Name
}

func (s DomainKitchen) RootFolderPath() string {

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

func (s DomainKitchen) RootTemplatesFolder() string {
	return fmt.Sprintf("./%s/web/templates", s.RootFolderPath())
}
