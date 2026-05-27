package maths

import (
	"math"

	"github.com/1Vewton/vocabulary_autorecord/data_management/config"
)

// Get difference of days between two unix timestamps
func GetDays(start int, end int) float64 {
	second_difference := end - start
	return float64(second_difference) / (86400.0)
}

// Get possibility of studied using BKT
func GetPossibility(
	studied_possibility float64,
	is_correct bool) float64 {
	var result_possibility float64
	// predicting possibility of getting correct
	var p_correct float64 = (studied_possibility*(1-config.GetpS()) +
		(1-studied_possibility)*config.GetpG())
	var p_answer float64
	// Based on wether it is correct or not
	if is_correct {
		p_answer = (studied_possibility * (1 - config.GetpS())) / p_correct
	} else {
		p_answer = (studied_possibility * config.GetpS()) / (1 - p_correct)
	}
	result_possibility = p_answer + (1-p_answer)*config.GetpT()
	return result_possibility
}

// Time decay
func PossibilityDecay(possibility float64, days float64) float64 {
	decay_constant := math.Exp(-config.GetLambda() * days)
	return max(possibility*decay_constant, config.GetpL0())
}
