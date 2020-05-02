package handler

import (
	"fmt"

	"github.com/gobuffalo/packr/v2"
	"github.com/lbernardo/odin/internal"
	"github.com/lbernardo/odin/pkg/models"
	"github.com/spf13/viper"
)

type CreateCmd struct {
	box     *packr.Box
	odinDir string
	config  models.OdinConfig
	module  models.Module
}

func NewCreateCmd(box *packr.Box) *CreateCmd {
	var configModels models.OdinConfig
	viper.Unmarshal(&configModels)

	return &CreateCmd{
		box:     box,
		odinDir: viper.GetString("ODIN_DIR"),
		config:  configModels,
		module:  internal.ReadYaml(viper.GetString("ODIN_DIR") + "/modules/" + configModels.Config.Default),
	}
}

func (cc *CreateCmd) CreateProject(name string) {
	internal.CreatePaths("", []string{name})
	internal.CreatePaths(name, cc.module.Create.Directories)
	fmt.Println(cc.module.Create.Files)
}
