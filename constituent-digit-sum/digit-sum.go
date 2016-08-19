package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func sumDigits(number string) int {
	n, _ := strconv.Atoi(number)
	sum := 0

	for i := len(number) - 1; i >= 0; i-- {
		m := int(math.Floor((float64(n) / math.Pow10(i))))
		sum += m
		n -= int(float64(m) * math.Pow10(i))
	}

	return sum
}

func main() {
	argc, argv := len(os.Args), os.Args

	if argc < 2 {
		fmt.Fprintf(os.Stderr, "I need at least one argument.\n")
		os.Exit(1)
	}

	fmt.Println(sumDigits(argv[1]))
}
