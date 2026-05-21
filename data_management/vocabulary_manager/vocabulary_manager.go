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

// Read vocabulary list from file
func ReadVocabularyList() (
	schema_content_bytes []byte,
	content_bytes []byte,
	schema_content_string string,
	content_string string,
	Error error,
) {
	// Check whether the file exists or not
	_, err := os.Stat(config.Settings.VocabListPath)
	if os.IsNotExist(err) {
		return nil, nil, "", "", errors.New("Vocabulary list file does not exist")
	}
	// Read file
	content, err := os.ReadFile(config.Settings.VocabListPath)
	if err != nil {
		return nil, nil, "", "", err
	}
	string_content := string(content)
	// Read schema
	schema_content, err := os.ReadFile(config.Settings.VocabSchemaPath)
	string_schema_content := string(schema_content)
	if err != nil {
		return nil, nil, "", "", err
	}
	return schema_content, content, string_schema_content, string_content, nil
}

// Add vocabularies to the vocabulary list
func AddVocabularies(vocab_list_from_file map[string]string) (Error error) {
	fmt.Println("Start adding vocabulary...")
	var string_schema_content string
	var string_content string
	var content []byte
	// Read schema and content
	_, content, string_schema_content, string_content, err := ReadVocabularyList()
	if err != nil {
		return err
	}
	// validate json
	result, err := json_validator.Validate(string_schema_content, string_content)
	fmt.Println("Finish validating json")
	if result {
		var vocabulary_list VocabularyList
		err = json.Unmarshal(content, &vocabulary_list)
		if err != nil {
			return err
		}
		// Add new vocabularies
		for vocab, def := range vocab_list_from_file {
			fmt.Printf("\033[32mAdd new vocabulary: %s\033[0m", vocab)
			fmt.Println()
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
				new_vocab.StudiedPossibility = config.GetpL0()
				vocabulary_list.Data = append(vocabulary_list.Data, new_vocab)
			} else {
				fmt.Println("\033[31mThis vocabulary already exists!\033[0m")
				fmt.Println()
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
	} else {
		return errors.New("Invalid json format")
	}
	return nil
}

// Add single vocabulary to the vocabulary list
func AddVocabulary(vocab string, def string) (Error error) {
	fmt.Println("Start adding vocabulary...")
	var string_schema_content string
	var string_content string
	var content []byte
	// Read schema and content
	_, content, string_schema_content, string_content, err := ReadVocabularyList()
	if err != nil {
		return err
	}
	// validate json
	result, err := json_validator.Validate(string_schema_content, string_content)
	fmt.Println("Finish validating json")
	if result {
		var vocabulary_list VocabularyList
		err = json.Unmarshal(content, &vocabulary_list)
		if err != nil {
			return err
		}
		// Add new vocabulary
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
			new_vocab.StudiedPossibility = config.GetpL0()
			vocabulary_list.Data = append(vocabulary_list.Data, new_vocab)
		} else {
			fmt.Println("\033[31mThis vocabulary already exists!\033[0m")
			fmt.Println()
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
	} else {
		return errors.New("Invalid json format")
	}
	return nil
}
