package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputText := "2 3 10"

	inputArray := strings.Split(inputText, " ")

	fizz, _ := strconv.Atoi(inputArray[0])
	buzz, _ := strconv.Atoi(inputArray[1])
	n, _ := strconv.Atoi(inputArray[2])

	for i := 1; i <= n; i++ {
		if i%fizz == 0 && i%buzz == 0 {
			fmt.Printf("FB")
		} else if i%fizz == 0 {
			fmt.Printf("F")
		} else if i%buzz == 0 {
			fmt.Printf("B")
		} else {
			fmt.Print(i)
		}

		if i < n {
			fmt.Printf(" ")
		} else {
			fmt.Printf("\n")
		}
	}
}
