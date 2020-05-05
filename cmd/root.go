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
	"fmt"
	"os"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var Box = packr.New("resources", "../resources")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "odin",
	Short: "Create your application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	initConfig()
	LoadCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	homeDir := home + "/.odin"

	// Search config in home directory with name ".odin" (without extension).
	viper.AddConfigPath(homeDir)
	viper.SetConfigFile(homeDir + "/config.yml")
	viper.Set("ODIN_DIR", homeDir)

	viper.AutomaticEnv() // read in environment variables that match

	viper.ReadInConfig()

}
