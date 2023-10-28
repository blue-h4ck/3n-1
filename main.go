package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	for {
		var input int
		fmt.Print("Enter some positive number >> ")
		var inputStr string
		_, err := fmt.Scanln(&inputStr)
		if err != nil {
			fmt.Println("\nTry Again Error >> ", err)
			continue
		}
		inputStr = strings.TrimSpace(inputStr)
		if inputStr == "" {
			fmt.Println("\nTry Again Error >> Expected an integer")
			continue
		}
		_, err = fmt.Sscan(inputStr, &input)
		if input > 0 {
		} else {
			continue
		}
		if err != nil {
			fmt.Println("\nTry Again Error >> Expected an integer")
			continue
		}
		// Making the 3n+1 rules
		var count int
		count = 1
		for {
			if input == 1 && count == 2 {
				count = 0
				fmt.Println("You got 1 again")
				break
			} else if input == 1 {
				count += 1
			}
			if input%2 == 0 {
				input = input / 2
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Number: ", input)
				continue
			} else {
				input = 3*input + 1
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Number: ", input)
				continue
			}
		}
	}
}
