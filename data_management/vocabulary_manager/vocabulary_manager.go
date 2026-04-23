package vocabulary_manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
	"github.com/1Vewton/vocabulary_autorecord/utils/json_validator"
)

// Vocabulary struct
type Vocabulary struct {
	Word               string
	Definition         string
	StudiedPossibility float64
}

type VocabularyList struct {
	LastUpdateTime int
	Data           []Vocabulary
}

// Initializes vocabulary list by checking whether the file exists or not
func InitializeVocabularyManager() (Error error) {
	_, err := os.Stat(config.Settings.VocabListPath)
	if os.IsNotExist(err) {
		_, err = os.Create(config.Settings.VocabListPath)
		if err != nil {
			return err
		}
		// Initialize vocabulary list
		var vocabulary_list_init VocabularyList
		vocabulary_list_init.LastUpdateTime = int(time.Now().Unix())
		vocabulary_list_init.Data = make([]Vocabulary, 0)
		// Write it to file
		bytes, err := json.MarshalIndent(vocabulary_list_init, "", "  ")
		if err != nil {
			return err
		}
		err = os.WriteFile(config.Settings.VocabListPath, bytes, 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

// Add vocabularies to the vocabulary list
func AddVocabularies(vocab_list_from_file map[string]string) (Error error) {
	fmt.Println("Start add vocabulary...")
	// Check whether the file exists or not
	_, err := os.Stat(config.Settings.VocabListPath)
	if os.IsNotExist(err) {
		return errors.New("Vocabulary list file does not exist")
	}
	// Read file
	content, err := os.ReadFile(config.Settings.VocabListPath)
	if err != nil {
		return err
	}
	string_content := string(content)
	// Read schema
	schema_content, err := os.ReadFile(config.Settings.ConfigSchemaPath)
	string_schema_content := string(schema_content)
	// validate json
	result, err := json_validator.Validate(string_schema_content, string_content)
	if result {
		var vocabulary_list VocabularyList
		err = json.Unmarshal(content, &vocabulary_list)
		if err != nil {
			return err
		}
		// Add new vocabularies
		for vocab, def := range vocab_list_from_file {
			fmt.Printf("\033[32mAdd new vocabulary: %s\033[0m", vocab)
			vocabulary_already_exists := false
			for _, v := range vocabulary_list.Data {
				if v.Word == vocab {
					vocabulary_already_exists = true
					break
				}
			}
			if !vocabulary_already_exists {
				var new_vocab Vocabulary
				new_vocab.Word = vocab
				new_vocab.Definition = def
				new_vocab.StudiedPossibility = 0.0
				vocabulary_list.Data = append(vocabulary_list.Data, new_vocab)
			} else {
				fmt.Println("\033[31mThis vocabulary already exists!\033[0m")
			}
		}
		// turn it to bytes and store it in file
		bytes, err := json.MarshalIndent(vocabulary_list, "", "  ")
		if err != nil {
			return err
		}
		err = os.WriteFile(config.Settings.VocabListPath, bytes, 0666)
		if err != nil {
			return err
		}
		fmt.Println("Add vocabulary successfully!")
	}
	return nil
}
