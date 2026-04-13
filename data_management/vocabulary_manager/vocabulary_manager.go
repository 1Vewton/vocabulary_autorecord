package vocabulary_manager

import (
	"os"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
)

// Initializes vocabulary list by checking whether the file exists or not
func InitializeVocabularyManager() (Error error) {
	_, err := os.Stat(config.Settings.VocabListPath)
	if os.IsNotExist(err) {
		_, err = os.Create(config.Settings.VocabListPath)
		if err != nil {
			return err
		}
	}
	return nil
}
