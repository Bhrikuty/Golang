package main

import "fmt"

type contact struct {
	gmail string
	zip   int
}

type person struct {
	fname string
	lname string
	contact
}

func main() {
	a := person{
		fname: "hello",
		lname: "world",
		contact: contact{
			gmail: "abc@gmail.com",
			zip:   123,
		},
	}
	aPointer := &a               // made pointer a which refers to memory address where this struct is pointing
	aPointer.updatename("jimmy") //a.updatename("jimmy") <- this was also correct as we can pass person type to pointer to person (no need to use line 25 in this case)
	a.print()

}

func (p *person) updatename(newfname string) { // before pointers asterick was nt there
	(*p).fname = newfname // p.fname = newfname <- here GO makes a copy of "a" as p and updates that p =this is called call by value, so we need to make a pointer

}
func (p person) print() {
	fmt.Printf("%+v", p)
}
