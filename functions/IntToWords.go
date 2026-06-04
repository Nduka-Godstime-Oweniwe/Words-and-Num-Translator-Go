package translator

import (
	"strconv"
	"strings"
)

func Type(n int) string {
	if n/1000000000000 != 0 {
		return "trillion"
	}

	if n/1000000000 != 0 {
		return "billion"
	}

	if n/1000000 != 0 {
		return "million"
	}

	if n/1000 != 0 {
		return "thousand"
	}
	return ""
}

func ConvertToWord(num int) string {
	first20 := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen",
		"twenty",
	}

	for i := 0; i < len(first20); i++ {
		if num == i+1 {
			return first20[i]
		}
	}

	next := []string{
		"thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety",
	}

	for k := 0; k < len(next); k++ {
		if num == (k+3)*10 {
			return next[k]
		}
	}

	// if str == "thousand" {
	// 	return 1000
	// }

	// if str == "million" {
	// 	return 1000000
	// }

	// if str == "billion" {
	// 	return 1000000000
	// }

	// if str == "trillion" {
	// 	return 1000000000000
	// }

	return ""
}
func GroupOfThreeStr(str string) string {
	first, _ := strconv.Atoi(string(str[0]))
	First := ""
	if first != 0 {
		First = ConvertToWord(first) + " hundred"
		if str[1:] != "00" {
			First += " and"
		} else if str[1:] == "00" {
			return First
		}
	}

	second, _ := strconv.Atoi(string(str[1]))
	third, _ := strconv.Atoi(string(str[2]))
	Second := ""
	num := second*10 + third
	if second != 0 {
		if num <= 20 || num%10 == 0 {
			Second = ConvertToWord(num)
		} else {
			Second = ConvertToWord(second*10) + " " + ConvertToWord(third)
		}
	} else {
		Second = ConvertToWord(third)
	}

	if (first == 0 && second == 0) || first == 0 {
		return Second
	}
	return First + " " + Second
}

func ConvertGroupIntToStr(n int) string {
	if n > 99 {
		return strconv.Itoa(n)
	}

	if n > 9 {
		return "0" + strconv.Itoa(n)
	}

	return "00" + strconv.Itoa(n)
}

func TranslateToWords(n int) string {
	str := strconv.Itoa(n)
	if (len(str)+1)%3 == 0 {
		str = "0" + str
	} else if (len(str)+2)%3 == 0 {
		str = "00" + str
	}
	groupOfThrees := []string{}
	group := ""
	for i, v := range str {
		group += string(v)
		if (i+1)%3 == 0 {
			groupOfThrees = append(groupOfThrees, group)
			group = ""
		}

	}

	names := []string{
		"trillion",
		"billion",
		"million",
		"thousand",
		"",
	}
	name := []string{}
	value := []string{}
	for i := 0; i < len(groupOfThrees); i++ {
		if len(groupOfThrees) == 5 {
			name = append(name, names[i])
			value = append(value, groupOfThrees[i])

		} else if len(groupOfThrees) == 4 {
			name = append(name, names[i+1])
			value = append(value, groupOfThrees[i])

		} else if len(groupOfThrees) == 3 {
			name = append(name, names[i+2])
			value = append(value, groupOfThrees[i])

		} else if len(groupOfThrees) == 2 {
			name = append(name, names[i+3])
			value = append(value, groupOfThrees[i])

		} else {
			name = append(name, names[i+4])
			value = append(value, groupOfThrees[i])

		}
	}

	result := ""

	for i := 0; i < len(name); i++ {
		if i == len(name)-1 {
			if string(value[0][0]) == "0" {
				result += "and "
			}
		}
		if value[i] == "000" {
			continue
		}
		result += GroupOfThreeStr(value[i]) + " " + name[i] + " "
		// if value != "" {
		// 	string += " " + value
		// }
	}
	result = strings.TrimSpace(result)
	return result

}
