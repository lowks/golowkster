package main

import "fmt"
import "path/filepath"

func PrintSomeStuff(input string) string {
	return fmt.Sprintf("hello %s", input)
}

func addTwoNumbers(x int, y int) int {
    return x + y
}

func JoinPath(first string, second string) string {
    return fmt.Sprintf(filepath.Join(first, second))
}
