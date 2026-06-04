package translator

import "strings"

func ValidStr(str string) bool {
	if str == "" {
		return false
	}
	slice := strings.Fields(str)
	for i := 0; i < len(slice); i++ {
		if !IsValidWord(slice[i]) {
			return false
		}
	}

	if slice[0] == "trillion" || slice[0] == "billion" || slice[0] == "million" {
		return false
	}
	return true
}
