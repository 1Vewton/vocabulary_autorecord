package vocabulary_manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
	"github.com/1Vewton/vocabulary_autorecord/utils/confirmation_interface"
	"github.com/1Vewton/vocabulary_autorecord/utils/json_validator"
	"github.com/1Vewton/vocabulary_autorecord/utils/maths"
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

// Sort interface
func (a VocabularyList) Len() int {
	return len(a.Data)
}
func (a VocabularyList) Swap(i, j int) {
	a.Data[i], a.Data[j] = a.Data[j], a.Data[i]
}
func (a VocabularyList) Less(i, j int) bool {
	return a.Data[i].StudiedPossibility < a.Data[j].StudiedPossibility
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

// Get vocabulary list
func GetVocabularyList() (VocabularyList, error) {
	var vocabulary_list VocabularyList
	// Read schema and content
	_, content, string_schema_content, string_content, err := ReadVocabularyList()
	if err != nil {
		return vocabulary_list, err
	}
	// validate json
	result, err := json_validator.Validate(string_schema_content, string_content)
	if result {
		err = json.Unmarshal(content, &vocabulary_list)
		if err != nil {
			return vocabulary_list, err
		}
		// Sort
		sort.Sort(vocabulary_list)
		// Update the studied possibility of each vocabulary
		fmt.Printf("current time: %d\n", int(time.Now().Unix()))
		days := maths.GetDays(vocabulary_list.LastUpdateTime, int(time.Now().Unix()))
		fmt.Printf("%f days since last update\n", days)
		for i := 0; i < len(vocabulary_list.Data); i++ {
			original_posibility := vocabulary_list.Data[i].StudiedPossibility
			vocabulary_list.Data[i].StudiedPossibility = maths.PossibilityDecay(original_posibility, days)
			fmt.Printf("%f -> %f\n", original_posibility, vocabulary_list.Data[i].StudiedPossibility)
		}
		return vocabulary_list, nil
	} else {
		return vocabulary_list, errors.New("Invalid json format")
	}
}

// Delete vocabulary from the vocabulary list
func DeleteVocabulary(vocab string) (bool, error) {
	var delete_result bool = false
	fmt.Println("Start deleting vocabulary...")
	// Read schema and content
	_, content, string_schema_content, string_content, err := ReadVocabularyList()
	if err != nil {
		return delete_result, err
	}
	// validate json
	result, err := json_validator.Validate(string_schema_content, string_content)
	fmt.Println("Finish validating json")
	if result {
		var vocabulary_list VocabularyList
		err = json.Unmarshal(content, &vocabulary_list)
		if err != nil {
			return delete_result, err
		}
		// Delete vocabulary
		for i, v := range vocabulary_list.Data {
			if v.Word == vocab {
				vocabulary_list.Data = append(vocabulary_list.Data[:i], vocabulary_list.Data[i+1:]...)
				delete_result = true
				break
			}
		}
		// turn it to bytes and store it in file
		bytes, err := json.MarshalIndent(vocabulary_list, "", "  ")
		if err != nil {
			return delete_result, err
		}
		err = os.WriteFile(config.Settings.VocabListPath, bytes, 0666)
		if err != nil {
			return delete_result, err
		}
		fmt.Println("Add vocabulary successfully!")
	} else {
		return delete_result, errors.New("Invalid json format")
	}
	return delete_result, nil
}

// Vocabulary management
func VocabularyManagement() (Error error) {
	// Initialize vocabulary list
	vocabulary_list, err := GetVocabularyList()
	if err != nil {
		fmt.Printf("\033[31mError: %s\033[0m", err)
		return err
	}
	for vocab := range vocabulary_list.Data {
		fmt.Printf("Word: %s", vocabulary_list.Data[vocab].Word)
		fmt.Println()
		fmt.Printf("Definition: %s", vocabulary_list.Data[vocab].Definition)
		fmt.Println()
		fmt.Println()
	}
	var delete_vocab bool = true
	for delete_vocab {
		delete_vocab = confirmation_interface.ConfirmationInterface("Do you want to delete a vocabulary?", false)
		if delete_vocab {
			var vocab_deleted bool
			var input string
			fmt.Println("Type the word of the vocabulary to delete:")
			fmt.Scan(&input)
			vocab_deleted, err = DeleteVocabulary(input)
			if err != nil {
				fmt.Printf("\033[31mError: %s\033[0m", err)
				fmt.Println()
				return err
			}
			if vocab_deleted {
				fmt.Printf("\033[32mVocabulary %s has been deleted successfully!\033[0m", vocabulary_list.Data[0].Word)
				fmt.Println()
			} else {
				fmt.Println("\033[31mError: Vocabulary not found!\033[0m")
			}
		} else {
			fmt.Println("No vocabulary to delete.")
		}
	}
	return nil
}

func SaveVocabularyList(updated_vocabulary_list VocabularyList) error {
	vocabulary_list := updated_vocabulary_list
	// Update time
	vocabulary_list.LastUpdateTime = int(time.Now().Unix())
	// turn it to bytes and store it in file
	bytes, err := json.MarshalIndent(vocabulary_list, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(config.Settings.VocabListPath, bytes, 0666)
	if err != nil {
		return err
	}
	fmt.Println("Vocabulary saving successfully!")
	return nil
}
