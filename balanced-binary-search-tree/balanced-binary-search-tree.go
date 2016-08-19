package main

import (
	"fmt"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
	bal   int
}
