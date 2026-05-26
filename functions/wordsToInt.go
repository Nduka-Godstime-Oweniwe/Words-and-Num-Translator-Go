package translator

func IsValidWord(str string) bool {
	Words := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen",
		"twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninty",
		"hundred", "thousand", "million",
		"billion", "trillion", "and"}
	for k := 0; k < len(Words); k++ {
		if str == Words[k] {
			return true
		}
	}
	return false
}

func IsValidInput(slice []string) bool {
	for k := 0; k < len(slice); k++ {
		if !IsValidWord(slice[k]) {
			return false
		}
	}
	return true
}

func ConvertToInt(str string) int {
	first20 := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen",
		"twenty",
	}

	for i := 0; i < len(first20); i++ {
		if str == first20[i] {
			return i + 1
		}
	}

	next := []string{
		"thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninty",
		"hundred",
	}

	for k := 0; k < len(next); k++ {
		if str == next[k] {
			return (k + 3) * 10
		}
	}

	if str == "thousand" {
		return 1000
	}

	if str == "million" {
		return 1000000
	}

	if str == "billion" {
		return 1000000000
	}

	if str == "trillion" {
		return 1000000000000
	}

	return 0
}

func SliceOfInt(str []string) []int {
	slice := []int{}
	for k := 0; k < len(str); k++ {
		if str[k] == "and" {
			continue
		}
		slice = append(slice,
			ConvertToInt(str[k]))
	}

	return slice
}

func IsTens(n int) bool {
	if n == 100 || n == 1000 || n == 1000000 {
		return true
	}
	if n == 1000000000 || n == 1000000000000 {
		return true
	}

	return false
}
