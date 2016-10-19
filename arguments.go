package main

import "fmt"
import "os"

func main() {
    args := os.Args
    arg1 := os.Args[1:]

    fmt.Println(args)
    fmt.Println(arg1)
}

