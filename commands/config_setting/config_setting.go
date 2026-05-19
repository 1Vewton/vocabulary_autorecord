package config_setting

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/basic_config"
	"github.com/spf13/cobra"
)

// Config Setting command
var ConfigSettingCmd = &cobra.Command{
	Use:   "configSetting",
	Short: "Changing the configuration of the program",
	Long:  "Changing the configuration of the program and save it to the config file.",
	Run: func(cmd *cobra.Command, args []string) {
		Error := basic_config.ChangeConfig()
		if Error != nil {
			fmt.Printf("\033[31mConfig changing failed due to %s\033[0m", Error)
			fmt.Println()
			return
		}
		fmt.Println("\033[32mConfig changing successful\033[0m")
	},
}
