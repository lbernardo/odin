package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/lbernardo/odin/pkg/models"
)

func WriteFile(project, path, content string, force bool) error {
	p := path
	if project != "" {
		p = fmt.Sprintf("%v/%v", project, path)
	}

	if !force {
		if _, err := os.Stat(p); os.IsExist(err) {
			return nil
		}
	}

	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write([]byte(content))

	fmt.Println("Created file", p)

	return nil
}

func CreatePaths(project string, paths []string) error {
	for _, p := range paths {
		path := p
		if project != "" {
			path = fmt.Sprintf("%v/%v", project, p)
		}

		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
		fmt.Println("Created ", path)
	}
	return nil
}

func CopyFile(o, d string, module models.Module, box *packr.Box) {
	var content string
	if module.Resource == "./" {
		content, _ = box.FindString(o)
	} else {
		dat, err := ioutil.ReadFile(o)
		if err != nil {
			fmt.Println("Error to read ", o)
		}
	}
}
