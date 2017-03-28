package main

import "fmt"

var p = fmt.Println

func main() {
	my_map := make(map[string]string)
	my_map["idiot"] = "zhuchini"
	p(my_map["idiot"])
	p("The length of my_app is:", len(my_map))
	delete(my_map, "idiot")
	p("The length of my_app after delete is:", len(my_map))	
}