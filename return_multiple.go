package main

import "fmt"

func return_multiple(input int, a_string string) (int, string) {
	return input, a_string
}

func main() {
    fmt.Println(return_multiple(5, "hoho"))
}
