package root

import (
	"github.com/1Vewton/vocabulary_autorecord/commands/config_setting"
	"github.com/1Vewton/vocabulary_autorecord/commands/read_file"
	"github.com/spf13/cobra"
)

// Root command
var rootCommand = &cobra.Command{
	Use:   "autorecord",
	Short: "A tool to automatically record vocabulary",
	Long:  "Autorecord is a tool to automatically record vocabulary. It can be used to record vocabularies in xlsx files and test the users about the vocabularies",
}

// Execute the root command
func Execute() error {
	rootCommand.AddCommand(read_file.ReadFileCMD)
	rootCommand.AddCommand(config_setting.ConfigSettingCmd)
	err := rootCommand.Execute()
	return err
}
