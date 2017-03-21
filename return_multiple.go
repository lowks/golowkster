package main

import "fmt"

func return_multiple(input int, a_string string) (input_int int, b_string string) {
	input_int = input
	b_string = a_string
	return
}

func main() {
    fmt.Println(return_multiple(5, "hoho"))
}
