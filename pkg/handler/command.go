package handler

import (
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/lbernardo/odin/internal"
	"github.com/lbernardo/odin/pkg/models"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Command struct {
	box     *packr.Box
	odinDir string
	config  models.OdinConfig
	module  models.Module
	cmd     models.Command
	flags   *pflag.FlagSet
	project string
}

func NewCommand(box *packr.Box, cmd models.Command, flags *pflag.FlagSet) {
	var configModels models.OdinConfig
	viper.Unmarshal(&configModels)

	project, _ := flags.GetString("project")

	if project == "" {
		project = "./"
	}

	cc := Command{
		box:     box,
		odinDir: viper.GetString("ODIN_DIR"),
		config:  configModels,
		module:  internal.ReadYaml(viper.GetString("ODIN_DIR") + "/modules/" + configModels.Config.Default),
		cmd:     cmd,
		flags:   flags,
		project: project,
	}

	// Set project
	viper.Set("ODIN_PROJECT", project)

	cc.executeCmd()
}

func (cc *Command) executeCmd() {
	for _, f := range cc.cmd.Files {
		p := strings.Split(cc.replaceVars(f), ":")
		internal.CopyFile(p[0], p[1], cc.module, cc.box, cc.getArgs())
	}
	for _, d := range cc.cmd.Directories {
		internal.CreatePaths(cc.project, []string{d})
	}
}

func (cc *Command) replaceVars(content string) string {
	for _, v := range cc.cmd.Args {
		value, _ := cc.flags.GetString(v.Name)
		content = strings.ReplaceAll(content, "${"+v.Name+"}", strings.ToLower(value))
	}

	return content
}

func (cc *Command) getArgs() map[string]string {
	list := make(map[string]string, 0)
	for _, v := range cc.cmd.Args {
		value, _ := cc.flags.GetString(v.Name)
		list[v.Name] = value
	}
	return list
}
