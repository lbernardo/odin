package handler

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/lbernardo/odin/internal"
	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/fmt/errors"
)

type ConfigCmd struct {
	box     *packr.Box
	odinDir string
}

func NewConfigCmd(box *packr.Box) *ConfigCmd {
	return &ConfigCmd{
		box:     box,
		odinDir: viper.GetString("ODIN_DIR"),
	}

}

func (cc *ConfigCmd) NewModule(name string) {

	module, err := cc.box.FindString("modules/module.template.yml")
	if err != nil {
		panic(err)
	}

	if err := internal.WriteFile("", cc.odinDir+"/modules/"+name+".yml", module, false); err != nil {
		panic(err)
	}

	fmt.Println("Edit " + cc.odinDir + "/modules/" + name + ".yml")
}

func (cc *ConfigCmd) NewDefault(name string) {

	configYaml, err := cc.box.FindString("config.yml")
	if err != nil {
		panic(err)
	}

	configYaml = strings.ReplaceAll(configYaml, "${default}", name+".yml")

	if err := internal.WriteFile("", cc.odinDir+"/config.yml", configYaml, true); err != nil {
		panic(errors.New("Error to created config.yml"))
	}
}
