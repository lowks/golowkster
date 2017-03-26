package main

import "fmt"
import "strings"

var p = fmt.Println

func main() {
	a := []string{"one", "two", "three"}
	b := strings.Join(a, ",")
	p(b)
	c := strings.Split(b, ",")
	p("%v", c)
	p("Contains", strings.Contains(b, "one"))
	var a_string string = " Hello World  "
	p("After Trimspace", strings.TrimSpace(a_string))
	var long_string string = `I am a very very very long
very very long string`
	var multi_line_string string = "I am a multiline" +
		"a multiline string which is very long" +
		"very long string"
	p("A very long string", long_string)
	p("Multiline string", multi_line_string)
	string_with_prefix := "prefixstring.tar"
	p("String with prefix", strings.HasPrefix(string_with_prefix, "prefixstring"))
	p("String with suffix", strings.HasSuffix(string_with_prefix, "tar"))
	p("Strings after split", strings.SplitAfter("abc,d,ef","abc"))
	p("Strings repeating", strings.Repeat("hoho",5))
	replacer := strings.NewReplacer("aaa","bbb")
	p("String replaced", replacer.Replace("aaa"))
}
