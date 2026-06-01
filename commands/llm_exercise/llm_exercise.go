package llm_exercise

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/exercise/llm"
	"github.com/spf13/cobra"
)

var vocabulary_num int

// LlmExerciseCommand
var LlmExerciseCMD = &cobra.Command{
	Use:   "llmExercise",
	Short: "Start LLM exercise",
	Long:  "Start an exercise that the definition will be checked by LLM over certain number of vocabularies",
	Run: func(cmd *cobra.Command, args []string) {
		err := llm.ExecuteExercise(vocabulary_num)
		if err != nil {
			fmt.Println("Error executing LLM exercise:", err)
			return
		}
	},
}

func init() {
	LlmExerciseCMD.Flags().IntVarP(&vocabulary_num, "num", "n", 10, "Number of vocabularies to be used in the exercise")
}
