package main

import ("fmt"
		"testing"
  		"github.com/stretchr/testify/assert")

func return_multiple(input int, a_string string) (input_int int, b_string string) {
	input_int = input
	b_string = a_string
	return
}

func TestReturnMultiple(t *testing.T) {
	assert.Equal(t, 123, 131, "they should be equal")
}

func main() {
    fmt.Println(return_multiple(5, "hoho"))
}
