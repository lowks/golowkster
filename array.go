package main
import "fmt"
import "strings"

var p = fmt.Println

func main() {
    a := []string {"one", "two", "three"}
    b := strings.Join(a, ",")
    p(b)
    c := strings.Split(b, ",")
    p("%v", c)
    p("Contains", strings.Contains(b, "one"))
}
