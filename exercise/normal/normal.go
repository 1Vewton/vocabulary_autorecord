package normal

import (
	"fmt"
	"time"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
	"github.com/1Vewton/vocabulary_autorecord/utils/clear_screen"
	"github.com/1Vewton/vocabulary_autorecord/utils/confirmation_interface"
)

// Execute Exercise
func ExecuteExercise(exercise_num int) error {
	fmt.Println("Start Executing Excercise, mode: Normal")
	VocabList, err := vocabulary_manager.GetVocabularyList()
	if err != nil {
		fmt.Printf("\033[31mError: %s, The exercise would quite due to this error.\033[0m", err)
		return err
	}
	idx_range := min(exercise_num, len(VocabList.Data))
	for i := 0; i < idx_range; i++ {
		var user_familiar bool = false
		// Only if the word is not familiar, ask the user to revise it
		for !user_familiar {
			fmt.Println("Word: ", VocabList.Data[i].Word)
			var show_definition bool = false
			// Predict the possibility of studying according to BKT
			var p_correct float64 = (VocabList.Data[i].StudiedPossibility*(1-config.GetpS()) +
				(1-VocabList.Data[i].StudiedPossibility)*config.GetpG())
			for !show_definition {
				show_definition = confirmation_interface.ConfirmationInterface("Have you get the definition  of this word in your mind?", true)
			}
			fmt.Println("Definition: ", VocabList.Data[i].Definition)
			var is_correct bool = false
			is_correct = confirmation_interface.ConfirmationInterface("Is the definition in your mind correct?", true)
			// Diffferent possibilities of studied according to BKT
			var p_answer float64
			if is_correct {
				user_familiar = true
				p_answer = (VocabList.Data[i].StudiedPossibility * (1 - config.GetpS())) / p_correct
			} else {
				user_familiar = false
				p_answer = (VocabList.Data[i].StudiedPossibility * config.GetpS()) / (1 - p_correct)
			}
			VocabList.Data[i].StudiedPossibility = p_answer + (1-p_answer)*config.GetpT()
			fmt.Printf("\033[32mNew Studied Possibility: %f\n\033[0m", VocabList.Data[i].StudiedPossibility)
			time.Sleep(1 * time.Second)
			clear_screen.ClearScreen()
		}
	}
	// Save the vocab list to file
	err = vocabulary_manager.SaveVocabularyList(VocabList)
	if err != nil {
		return err
	}
	return nil
}
