package main

import "fmt"

func main() {

	card := deckfromfile()
	res := card.tostring()
	fmt.Println([]byte(res))
}
