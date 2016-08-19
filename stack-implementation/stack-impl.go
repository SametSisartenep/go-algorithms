package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	Items []int
}

func (s *Stack) Push(n int) {
	s.Items = append(s.Items, n)
}

func (s *Stack) Pop() int {
	n := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return n
}

func printStack(arr string) {
	var stack Stack
	itIsCandidate := true

	inputArray := strings.Split(arr, " ")
	for _, v := range inputArray {
		n, _ := strconv.Atoi(v)
		stack.Push(n)
	}

	for len(stack.Items) > 0 {
		if itIsCandidate {
			fmt.Print(stack.Pop())
			fmt.Print(" ")
		} else {
			stack.Pop()
		}
		itIsCandidate = !itIsCandidate
	}
	fmt.Printf("\n")
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		printStack(scanner.Text())
	}
}
