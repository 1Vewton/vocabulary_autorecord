package root

import (
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
	err := rootCommand.Execute()
	return err
}
