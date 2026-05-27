package translator

import "strconv"

func ValidateInt(str string) (int, bool) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, false
	}
	if num <= 0 {
		return 0, false
	}
	if num > 999999999999999 {
		return 0, false
	}
	return num, true
}
