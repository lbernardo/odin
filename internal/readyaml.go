package internal

import (
	"io/ioutil"

	"github.com/lbernardo/odin/pkg/models"
	"gopkg.in/yaml.v2"
)

func ReadYaml(file string) models.Module {
	var m models.Module
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(dat, &m)

	return m
}
