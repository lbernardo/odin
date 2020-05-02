package handler

import (
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/lbernardo/odin/internal"
	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/fmt/errors"
)

type StartCmd struct {
	box     *packr.Box
	odinDir string
}

func NewStartCmd(box *packr.Box) {
	startCmd := StartCmd{
		box:     box,
		odinDir: viper.GetString("ODIN_DIR"),
	}
	startCmd.createHome()
	startCmd.createConfigYaml()
	startCmd.createModuleDefault()
}

func (sc *StartCmd) createHome() {
	dirs := []string{
		sc.odinDir,
		sc.odinDir + "/modules",
	}
	if err := internal.CreatePaths("", dirs); err != nil {

	}
}

func (sc *StartCmd) createConfigYaml() {
	configYaml, err := sc.box.FindString("config.yml")
	if err != nil {
		panic(err)
	}

	configYaml = strings.ReplaceAll(configYaml, "${default}", "default.yml")

	if err := internal.WriteFile("", sc.odinDir+"/config.yml", configYaml, false); err != nil {
		panic(errors.New("Error to created config.yml"))
	}
}

func (sc *StartCmd) createModuleDefault() {
	configYaml, err := sc.box.FindString("modules/default.yml")
	if err != nil {
		panic(err)
	}

	if err := internal.WriteFile("", sc.odinDir+"/modules/default.yml", configYaml, false); err != nil {
		panic(err)
	}

}
