package translator

import (
	"strconv"
	"strings"
)

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
		"sixty", "seventy", "eighty", "ninety",
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

func TensValue(n int) int {
	str := strconv.Itoa(n)
	ans := ""
	for _, v := range str {
		if v == '0' {
			ans += string(v)
		}
	}

	ans = "1" + ans
	result, _ := strconv.Atoi(ans)
	if result == 10000 {
		return 1000
	}

	if result == 10000000 {
		return 1000000
	}

	if result == 10000000000 {
		return 1000000000
	}

	if result == 10000000000000 {
		return 1000000000000
	}
	return result
}

func PrintNumberWithComma(n int) string {
	str := strconv.Itoa(n)
	result := ""
	if len(str)%3 == 0 {
		for i, v := range str {
			result += string(v)
			if (i+1)%3 == 0 && i != len(str)-1 {
				result += ","
			}
		}
	}

	if (len(str)+1)%3 == 0 {
		for i, v := range str {
			result += string(v)
			// if i == 1 {
			// 	result += string(v)
			// 	continue
			// }
			if (i+2)%3 == 0 && i != len(str)-1 {
				result += ","
			}
		}
	}

	if (len(str)+2)%3 == 0 {
		for i, v := range str {
			result += string(v)
			// if i == 1 {
			// 	result += string(v)
			// 	continue
			// }
			if i%3 == 0 && i != len(str)-1 {
				result += ","
			}
		}
	}

	return result

}

func TranslateToInt(words string) (int, string) {
	// fmt.Println(str)
	slice := strings.Fields(words)
	str := SliceOfInt(slice)
	if len(str) == 1 && str[0] == 100 {
		return 100, "100"
	}
	// fmt.Println(str)
	ans := []int{}
	value := 0

	for i := 0; i < len(str); i++ {
		if IsTens(str[i]) {
			if i == 0 {
				value = 1 * str[i]
			} else {
				value = value * str[i]
			}
		} else if i != 0 && IsTens(str[i-1]) {

			ans = append(ans, value)
			value = 0
			value += str[i]

		} else {
			value += str[i]

		}
	}

	if value != 0 {
		ans = append(ans, value)
	}
	// fmt.Println(ans)
	if len(ans) == 1 {
		return ans[0], PrintNumberWithComma(ans[0])
	}
	final := 0
	for i := 1; i < len(ans); i++ {
		if ans[i] > ans[i-1] {
			// fmt.Println(TensValue(ans[i]))
			k := ans[i-1]*TensValue(ans[i]) + ans[i]
			final += k
		} else if i == 1 {
			final += ans[i-1]
			if i != len(ans)-1 {
				if ans[i+1] > ans[i] {
					continue
				}
			}
			final += ans[i]

		} else if i != len(ans)-1 {
			if ans[i+1] > ans[i] {
				continue
			} else {
				final += ans[i]

			}
		} else {
			final += ans[i]
		}
		// } else {
		// 	final += ans[i]
		// }
	}
	return final, PrintNumberWithComma(final)
}
