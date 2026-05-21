package add_vocabulary

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
	"github.com/spf13/cobra"
)

var vocabulary string
var definition string

// Command to add a single vocabulary to the vocabulary list.
var AddVocabularyCMD = &cobra.Command{
	Use:   "addSingleVocabulary",
	Short: "The command to add a single vocabulary to the vocabulary list.",
	Long:  "The command to add a single vocabulary to the vocabulary list. The vocabulary and its definition should be provided as arguments.",
	Run: func(cmd *cobra.Command, args []string) {
		err := vocabulary_manager.AddVocabulary(vocabulary, definition)
		if err != nil {
			fmt.Printf("\033[31mVocabulary adding failed due to %s\033[0m", err)
		} else {
			fmt.Println("\033[32mVocabulary adding successful\033[0m")
		}
	},
}

// Init: Bind flag to command
func init() {
	AddVocabularyCMD.Flags().StringVarP(&vocabulary, "vocabulary", "v", "", "The vocabulary to be added")
	AddVocabularyCMD.Flags().StringVarP(&definition, "definition", "d", "", "The definition of the vocabulary to be added")
	AddVocabularyCMD.MarkFlagRequired("vocabulary")
	AddVocabularyCMD.MarkFlagRequired("definition")
}
