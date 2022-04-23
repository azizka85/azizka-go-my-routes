package helpers

import "strings"

func Condition(condition bool, trueStr string, falseStr string) string {
	if condition {
		return trueStr
	} else {
		return falseStr
	}
}

func StringToArray(param string) []string {
	array := strings.Split(param, ",")

	for i := 0; i < len(array); i++ {
		array[i] = strings.TrimSpace(array[i])
	}

	return array
}
