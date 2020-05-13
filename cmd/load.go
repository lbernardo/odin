/*
Copyright © 2020 Lucas Bernardo

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/lbernardo/odin/internal"
	"github.com/lbernardo/odin/pkg/handler"
	"github.com/lbernardo/odin/pkg/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadCommands() {

	if _, err := os.Stat(viper.GetString("ODIN_DIR")); !os.IsNotExist(err) {
		var configModels models.OdinConfig
		viper.Unmarshal(&configModels)

		yml := internal.ReadYaml(viper.GetString("ODIN_DIR") + "/modules/" + configModels.Config.Default)
		var c *cobra.Command
		var a models.Args
		for _, cmd := range yml.Commands {
			cc := func(cmd models.Command) *cobra.Command {
				c = &cobra.Command{
					Use:   cmd.Cmd,
					Short: cmd.Description,
					Run: func(c *cobra.Command, args []string) {
						handler.NewCommand(Box, cmd, c.Flags())
					},
				}
				for _, a = range cmd.Args {
					c.PersistentFlags().String(a.Name, a.Value, a.Description)
				}
				c.PersistentFlags().String("project", "", "Name of project")
				return c
			}
			rootCmd.AddCommand(cc(cmd))
		}
	}

}
