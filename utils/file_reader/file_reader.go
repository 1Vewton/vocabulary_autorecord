package file_reader

import (
	"errors"

	"github.com/1Vewton/vocabulary_autorecord/utils/error_interface"
	"github.com/xuri/excelize/v2"
)

// Excel Reader
func readExcel(
	filePath string,
	sheetName string,
	definitionRow string,
	vocabRow string,
) (map[string]string, error_interface.Error) {
	// Result storage
	result := make(map[string]string)
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
	// Get the index of the headers of rows
	columnMap := make(map[string]int)
	for i, header := range headers[0] {
		columnMap[header] = i
	}
	// Get the index for definition and vocab
	definitionIdx, ok := columnMap[definitionRow]
	if !ok {
		return result, error_interface.Error{
			IsError: true,
			Reason:  "Cannot find definition column",
			Err:     errors.New("Cannot find definition column"),
		}
	}
	vocabIdx, ok := columnMap[vocabRow]
	if !ok {
		return result, error_interface.Error{
			IsError: true,
			Reason:  "Cannot find vocab column",
			Err:     errors.New("Cannot find vocab column"),
		}
	}
	// Get data in these two rows
	vocabData := make(map[int]string)
	definitionData := make(map[int]string)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return result, error_interface.Error{
			IsError: true,
			Reason:  "Cannot get data rows",
			Err:     err,
		}
	}
	for i, row := range rows {
		if len(row) > definitionIdx && len(row) > vocabIdx {
			definitionData[i] = row[definitionIdx]
			vocabData[i] = row[vocabIdx]
		}
	}
	// Combine the data
	for i, vocab := range vocabData {
		definition, ok := definitionData[i]
		if ok {
			result[vocab] = definition
		}
	}
	// Return final result
	return result, error_interface.Error{
		IsError: false,
		Reason:  "",
		Err:     nil,
	}
}
