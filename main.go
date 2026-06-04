package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	translator "translator/functions"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	for {
		clearScreen()
		fmt.Println("1. Translate Numbers to Words")
		fmt.Println("2. Translate Words To Numbers")
		fmt.Println("3. Exit")
		option := translator.UserOption("Select An Option: ", 3)
		clearScreen()
		if option == 1 {
			for {

				num := translator.UserInput("Enter Number: ")
				number, valid := translator.ValidateInt(num)
				if !valid {
					fmt.Println("Invalid Number")
					continue
				}
				fmt.Println(translator.TranslateToWords(number))
				translator.UserInput("Press Enter To continue: ")
				break
			}
		} else if option == 2 {
			for {
				str := translator.UserInput("Enter number in words: ")
				valid := translator.ValidStr(str)
				if !valid {
					fmt.Println("Invalid Number in words")
					continue
				}
				num := translator.TranslateToInt(str)
				numWithComma := translator.PrintNumberWithComma(num)
				fmt.Println(numWithComma)
				translator.UserInput("Press Enter To continue: ")
				break
			}

		} else {
			break
		}

	}
}
