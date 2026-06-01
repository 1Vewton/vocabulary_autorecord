package result_data

import (
	"encoding/json"
)

type VocabularyCheckResult struct {
	correct bool
	notice  string
	tips    string
}

type VCRValidateResults struct {
	Data    VocabularyCheckResult
	Success bool
}

// Convert to VocabularyCheckResult
func ConvertVocabularyCheckResult(resultResp string) VCRValidateResults {
	var resultData VocabularyCheckResult
	var result VCRValidateResults
	err := json.Unmarshal([]byte(resultResp), &resultData)
	if err != nil {
		result.Success = false
		return result
	}
	result.Success = true
	result.Data = resultData
	return result
}

// is the result correct
func (vc VocabularyCheckResult) GetIsCorrect() bool {
	return vc.correct
}

// Get the notice
func (vc VocabularyCheckResult) GetNotice() string {
	return vc.notice
}
