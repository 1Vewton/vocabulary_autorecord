package read_file

import (
	"fmt"

	"github.com/1Vewton/vocabulary_autorecord/data_management/basic_config"
	"github.com/1Vewton/vocabulary_autorecord/data_management/vocabulary_manager"
	"github.com/1Vewton/vocabulary_autorecord/utils/file_reader"
	"github.com/spf13/cobra"
)

// path reading
var path string
var sheetName string

// Read file command
var ReadFileCMD = &cobra.Command{
	Use:   "readFile",
	Short: "Read vocab file content",
	Long:  "Read the vocabularies from the file and print them to the console",
	Run: func(cmd *cobra.Command, args []string) {
		// Read file content
		result, err_tmp := file_reader.ReadExcel(path, sheetName, basic_config.BasicConfig.DefinitionFieldName, basic_config.BasicConfig.VocabFieldName)
		if err_tmp.IsError {
			fmt.Printf("\033[31mFile reading failed due to %s\033[0m", err_tmp.Reason)
			fmt.Println()
			return
		}
		// Add it to vocab list
		err := vocabulary_manager.AddVocabularies(result)
		if err != nil {
			fmt.Printf("\033[31mVocabulary adding failed due to %s\033[0m", err)
			fmt.Println()
			return
		}
	},
}

// Init: Bind flag to command
func init() {
	ReadFileCMD.Flags().StringVarP(&path, "path", "p", "", "Path to the file to read")
	ReadFileCMD.Flags().StringVarP(&sheetName, "sheetName", "s", "", "Sheet name to read")
	ReadFileCMD.MarkFlagRequired("path")
	ReadFileCMD.MarkFlagRequired("sheetName")
}
