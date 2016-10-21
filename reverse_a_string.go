package main

import "fmt"

func reverse(input string) string {
    input_string := []rune(input)
    output_string := []rune{}

    for i := len(input_string) - 1; i >= 0; i -- {
        output_string = append(output_string, input_string[i])
    }

    return string(output_string)
}

func main() {
    string1 := "Low Kian Seong"
    reversed1 := reverse(string1)
    fmt.Println(reversed1)

    string2 := "gnoeS naiK woL"
    reversed2 := reverse(string2)
    fmt.Println(reversed2)
}
