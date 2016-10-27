package main

import "fmt"
import "flag"
import "strings"

var max *int = flag.Int("max", 0, "enter integer or bust!") 

func main() {
    flag.Parse()
    
    input_integer := *max
    
    for i := 1; i <= input_integer; i++ {
        process_fizz_buzz(i)
    }
}

func process_fizz_buzz(input int) {

    full_fizz_buzz := "CensofRocks" 
    // if input % 3 == 0 && input % 5 == 0 {
    //     fmt.Println(input, full_fizz_buzz)
    // } else if input % 3 == 0 {
    //     fmt.Println(input, strings.Replace(full_fizz_buzz, "buzz", "", -1))
    // } else if input % 5 == 0 {
    //     fmt.Println(input, strings.Replace(full_fizz_buzz,"fizz", "", -1))
    // } else {
    //     fmt.Println(input)
    // }
    switch {
        case input%15 == 0:
            fmt.Println(input, full_fizz_buzz)
        case input%3 == 0:
            fmt.Println(input, strings.Replace(full_fizz_buzz, "Rocks", "", -1))
        case input%5 == 0:
            fmt.Println(input, strings.Replace(full_fizz_buzz,"Censof", "", -1))
        default:
            fmt.Println(input)
    }
}