package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func ReturnMultiple(input int, astring string) (inputInt int, bstring string) {
// 	inputInt = input
// 	bstring = astring
// 	return
// }

func TestReturnMultiple(t *testing.T) {
	inputInteger, inputString := ReturnMultiple(10, "hello")
	assert.Equal(t, 10, inputInteger, "they should be equal")
	assert.Equal(t, "hello world", inputString, "they should be equal")
}

func TestHelloWorld(t *testing.T) {
	// fmt.Pr
	assert.Equal(t, JustReturnHelloWorld(), "HelloWorld", "These strings should be equal")
}

// func main() {
// 	fmt.Println(ReturnMultiple(5, "hoho"))
// }
