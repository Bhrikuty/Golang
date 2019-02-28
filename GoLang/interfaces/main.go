/* Bad Example- without inteface
package main

import "fmt"

type EngBot struct{}
type SpBot struct{}

func main() {

	eb := EngBot{}
	sb := SpBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb EngBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb SpBot) { //cannot do func over loading is cannot declare same func name again
	fmt.Println(sb.getGreeting())
}

func (EngBot) getGreeting() string { //here as we are not making use of eb and sb , we can remove it

	//custom logic for generating eng greeting
	return "hi there"
}

func (SpBot) getGreeting() string {

	//custom logic for generating spanish greeting
	return "hola"

}
*/

// good Example- with inteface

package main

import "fmt"

type bot interface {
	getGreeting() string //means if there is any other type in this program(i.e struct here) which uses getGreeting func and
	//returns a string, can now is a member of bot interface
}
type EngBot struct{}
type SpBot struct{}

func main() {

	eb := EngBot{}
	sb := SpBot{}
	printGreeting(eb) //'passed as struct' and 'accepted as interface' which is accepted acc to explaination above
	printGreeting(sb)
}

func printGreeting(b bot) { //since these types are member of bot interface and they can call the functions that accept something of type bot
	fmt.Println(b.getGreeting())
}

func (EngBot) getGreeting() string {

	//custom logic for generating eng greeting
	return "hi there"
}

func (SpBot) getGreeting() string {

	//custom logic for generating spanish greeting
	return "hola"

}
