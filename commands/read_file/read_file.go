package read_file

import (
	"github.com/spf13/cobra"
)

var readFileCMD = &cobra.Command{
	Use:   "readFile",
	Short: "Read vocab file content",
}
