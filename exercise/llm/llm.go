package llm

import (
	"errors"
	"fmt"
	"time"

	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
	"github.com/1Vewton/vocabulary_autorecord/utils/clear_screen"
	"github.com/1Vewton/vocabulary_autorecord/utils/llm/llm"
	"github.com/1Vewton/vocabulary_autorecord/utils/llm/prompt"
	"github.com/1Vewton/vocabulary_autorecord/utils/llm/result_data"
	"github.com/1Vewton/vocabulary_autorecord/utils/maths"
)

// Start the exercise
func ExecuteExercise(exercise_num int) error {
	fmt.Println("Start Executing Excercise, mode: LLM")
	VocabList, err := vocabulary_manager.GetVocabularyList()
	if err != nil {
		fmt.Printf("\033[31mError: %s, The exercise would quite due to this error.\033[0m", err)
		return err
	}
	exercise_length := min(len(VocabList.Data), exercise_num)
	for i := 0; i < exercise_length; i++ {
		var user_familiar bool = false
		for !user_familiar {
			fmt.Println("Word: ", VocabList.Data[i].Word)
			// Definition input
			fmt.Println("Input the definition of the word: ")
			var input_definition string
			fmt.Scan(&input_definition)
			// LLM Checking
			prompt := prompt.VocabularyCheckingPrompt(VocabList.Data[i].Word, VocabList.Data[i].Definition, input_definition)
			resp := make(chan string)
			err_chan := make(chan error)
			go llm.Request(prompt, resp, err_chan)
			// Process the response
			select {
			case response := <-resp:
				process_result := result_data.ConvertVocabularyCheckResult(response)
				if !process_result.Success {
					fmt.Println(response)
					return errors.New("Failed to process the respnse of the LLM")
				}
				is_correct := process_result.Data.GetIsCorrect()
				if is_correct {
					user_familiar = true
					fmt.Println("Correct! ")
				} else {
					user_familiar = false
					fmt.Println("Incorrect!  Please try again!")
					fmt.Println(process_result.Data.GetNotice())
				}
				VocabList.Data[i].StudiedPossibility = maths.GetPossibility(VocabList.Data[i].StudiedPossibility, is_correct)
				fmt.Printf("\033[32mNew Studied Possibility: %f\n\033[0m", VocabList.Data[i].StudiedPossibility)
				time.Sleep(1 * time.Second)
				clear_screen.ClearScreen()
			case err_res := <-err_chan:
				fmt.Printf("\033[31mError: %s, The exercise would quite due to this error.\033[0m", err_res)
				return err_res
			}
		}
	}
	return nil
}
