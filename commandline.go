package main
import "flag"
import "fmt"

func print_out_type(x interface{}) {
    switch x.(type) {
        case string:
             fmt.Println("It's a string")
        case int32:
             fmt.Println("It's an integer")
        default:
             fmt.Println("A default")
    }
}

func main() {
    wordPtr := flag.String("argument1", "foo", "a String")
    numPtr := flag.Int("argument2", 42, "an Integer")
    flag.Parse()
    fmt.Println("word: ", *wordPtr)
    fmt.Println("number: ", *numPtr)
    print_out_type(*wordPtr)
}
