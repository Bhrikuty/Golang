package main

import "fmt"

func main() {
	c := []string{"hello", newc()}
	c = append(c, "morning")
	fmt.Println(c)
}

func newc() string {
	return "good"
}
