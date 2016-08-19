package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func isJolly(items []int) bool {
	itemNo := items[0]

	if itemNo < 2 {
		return false
	}

	isJolly := true

	for i := 1; i < itemNo; i++ {
		diff := int(math.Abs(float64(items[i]) - float64(items[i+1])))
		if diff < 1 || diff >= itemNo {
			isJolly = false
		}
	}
	return isJolly
}

func printJollyness(input string) {
	inputArr := strings.Split(input, " ")
	var items []int

	for _, v := range inputArr {
		n, _ := strconv.Atoi(v)
		items = append(items, n)
	}

	if isJolly(items) {
		fmt.Println("Jolly")
	} else {
		fmt.Println("Not jolly")
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		printJollyness(scanner.Text())
	}
}
