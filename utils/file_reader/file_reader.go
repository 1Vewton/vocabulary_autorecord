package file_reader

import (
	"errors"

	"github.com/1Vewton/vocabulary_autorecord/utils/error_interface"
	"github.com/xuri/excelize/v2"
)

// Excel Reader
func readExcel(filePath string, sheetName string) (map[string]string, error_interface.Error) {
	// Result storage
	var result map[string]string
	// Open file
	f, err := excelize.OpenFile(filePath)
	// Check if the file is openned successfully
	if err != nil {
		return result, error_interface.Error{
			IsError: true,
			Reason:  "Failed to open file",
			Err:     err,
		}
	}
	// Get the header row
	headers, err := f.GetRows(sheetName)
	if err != nil {
		return result, error_interface.Error{
			IsError: true,
			Reason:  "Cannot get header row",
			Err:     err,
		}
	}
	if len(headers) == 0 {
		return result, error_interface.Error{
			IsError: true,
			Reason:  "No header row found",
			Err:     errors.New("There is no data in the table"),
		}
	}
	// Return final result
	return result, error_interface.Error{
		IsError: false,
		Reason:  "",
		Err:     nil,
	}
}
