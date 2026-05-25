package normal_exercise

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/exercise/normal"
	"github.com/spf13/cobra"
)

var vocabulary_num int

// Command to start normal exercise
var NormalExerciseCMD = &cobra.Command{
	Use:   "normalExercise",
	Short: "Start normal exercise",
	Long:  "Start a normal exercise over certain number of vocabularies",
	Run: func(cmd *cobra.Command, args []string) {
		err := normal.ExecuteExercise(vocabulary_num)
		if err != nil {
			fmt.Printf("\033[31mError executing normal exercise: %s\n\033[0m", err)
		}
	},
}

func init() {
	NormalExerciseCMD.Flags().IntVarP(&vocabulary_num, "num", "n", 10, "Number of vocabularies to be used in the exercise")
}
