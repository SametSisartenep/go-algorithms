package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func removeChars(str string) {
	strArr := strings.Split(str, ", ")
	str = strArr[0]
	chars := strArr[1]

	for i := 0; i < len(chars); i++ {
		str = strings.Replace(str, string(chars[i]), "", -1)
	}

	fmt.Println(str)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		removeChars(scanner.Text())
	}
}
