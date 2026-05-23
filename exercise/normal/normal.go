package normal

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
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

	}
	return nil
}
