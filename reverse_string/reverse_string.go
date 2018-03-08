package reverse_string

import (
	"fmt"
)

func reverse_string(input string) string {
	data := []rune(input)
	result := []rune{}
	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

func main() {
	input_string := "Hello World"
	reversed := reverse_string(input_string)
	fmt.Println(reversed)
}


