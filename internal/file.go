package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/lbernardo/odin/pkg/models"
	"github.com/spf13/viper"
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

func CopyFile(o, d string, module models.Module, box *packr.Box, args map[string]string) {
	var content string
	project := viper.GetString("ODIN_PROJECT")
	if project == "" {
		project = "./"
	}
	if module.Resource == "./" {
		o = strings.ReplaceAll(o, "${resource}", "")
		content, _ = box.FindString(o)
	} else {
		o = strings.ReplaceAll(o, "$resource", module.Resource)
		dat, err := ioutil.ReadFile(o)
		if err != nil {
			fmt.Println("Error to read ", o)
		}
		content = string(dat)
	}

	if args != nil {
		for n, v := range args {
			content = strings.ReplaceAll(content, "${"+n+"}", v)
		}
	}

	err := ioutil.WriteFile(project+"/"+d, []byte(content), 0755)
	if err != nil {
		fmt.Printf("Error create file %v/%v\n", project, d)
		return
	}
	fmt.Printf("Created  %v/%v\n", project, d)
}

func CreateConfigProject(name string, config models.ProjectConfig) {
	content := "pkg=" + config.Pkg
	ioutil.WriteFile(name+"/.odin", []byte(content), 0755)
}
