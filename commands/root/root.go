package root

import (
	"github.com/1Vewton/vocabulary_autorecord/commands/add_vocabulary"
	"github.com/1Vewton/vocabulary_autorecord/commands/config_setting"
	"github.com/1Vewton/vocabulary_autorecord/commands/llm_exercise"
	"github.com/1Vewton/vocabulary_autorecord/commands/normal_exercise"
	"github.com/1Vewton/vocabulary_autorecord/commands/read_file"
	"github.com/1Vewton/vocabulary_autorecord/commands/vocabulary_manager"
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
	rootCommand.AddCommand(add_vocabulary.AddVocabularyCMD)
	rootCommand.AddCommand(config_setting.ConfigSettingCmd)
	rootCommand.AddCommand(normal_exercise.NormalExerciseCMD)
	rootCommand.AddCommand(vocabulary_manager.VocabularyManagerCMD)
	rootCommand.AddCommand(llm_exercise.LlmExerciseCMD)
	err := rootCommand.Execute()
	return err
}
