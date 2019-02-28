package main

import "fmt"

func main() {

	// var colors map[string]string //empty map

	//colors:= make(map[string]string) // empty map,"make" is a built in func

	colors := map[string]string{ // key and val both r strings
		"red":   "ff00",
		"green": "ff11",
		"blue":  "f111",
	}
	// colors["white"] = "ffff" //append

	// delete(colors, "green") //delete any entry by mentioning key
	//fmt.Println(colors)
	print(colors)
}

func print(c map[string]string) {
	for k, v := range c {
		fmt.Println("hex code for", k, "is", v)
	}
}
