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
	"errors"

	"github.com/lbernardo/odin/pkg/handler"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure odin",
}

var configAddModule = &cobra.Command{
	Use:   "module [name]",
	Short: "Add new module template",
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a name module")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ch := handler.NewConfigCmd(Box)
		ch.NewModule(args[0])
	},
}

var configAddDefault = &cobra.Command{
	Use:   "default",
	Short: "Set default module",
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a name default module")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ch := handler.NewConfigCmd(Box)
		ch.NewDefault(args[0])
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configAddModule)
	configCmd.AddCommand(configAddDefault)

	// configAddModule.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	// configAddModule.MarkFlagRequired("region")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
