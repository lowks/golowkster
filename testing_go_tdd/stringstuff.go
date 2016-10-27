package main

import "fmt"
import "path/filepath"
import "strings"

func PrintSomeStuff(input string) string {
	return fmt.Sprintf("hello %s", input)
}

func addTwoNumbers(x int, y int) int {
    return x + y
}

func JoinPath(first string, second string) string {
    return fmt.Sprintf(filepath.Join(first, second))
}

func CapitalizeString(input string) string {
    return strings.ToUpper(input)
}

func CapitalizeFirstCharacter(input string) string {
    return strings.Title(input)
}
