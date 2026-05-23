package vocabulary_manager

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
	"github.com/spf13/cobra"
)

// Vocabulary Manager command
var VocabularyManagerCMD = &cobra.Command{
	Use:   "vocabularyManager",
	Short: "Manage the vocabularies",
	Long:  "Manage the vocabularies",
	Run: func(cmd *cobra.Command, args []string) {
		Error := vocabulary_manager.VocabularyManagement()
		if Error != nil {
			fmt.Printf("\033[31mError: %s, vocabulary management failed.\033[0m", Error)
			return
		}
	},
}
