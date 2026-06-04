package translator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func UserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func UserOption(prompt string, limit int) int {
	for {
		value, err := strconv.Atoi(UserInput(prompt))

		if err != nil || !(value >= 1 && value <= limit) {
			fmt.Println("Invalid Option!")
			continue
		}
		return value
	}

}
