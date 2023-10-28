/*
The 3n+1 problem, also known as the Collatz conjecture, is a simple mathematical problem that involves following a sequence of numbers. The problem is stated as follows:

Start with any positive integer n.
If n is even, divide it by 2.
If n is odd, multiply it by 3 and add 1.
Repeat the process with the new value obtained.
Continue this sequence until n becomes equal to 1.
The conjecture suggests that, regardless of the starting value of n, this sequence will always eventually reach the number 1. However, this statement has not been proven mathematically and remains an unsolved problem in mathematics. 
The 3n+1 problem is often used as a simple example in computer science and mathematics to illustrate concepts related to sequences, loops, and computational complexity.
*/

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
