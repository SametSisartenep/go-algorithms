package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseAndAdd(strno string, i int) string {
	strnor := reverseString(strno)
	n, _ := strconv.Atoi(strno)
	nr, _ := strconv.Atoi(strnor)
	iters := i

	sum := n + nr
	iters++
	sumstr := strconv.Itoa(sum)

	if sumstr == reverseString(sumstr) {
		return fmt.Sprintf("%d %d", iters, sum)
	} else {
		return reverseAndAdd(sumstr, iters)
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
		fmt.Println(reverseAndAdd(scanner.Text(), 0))
	}
}
