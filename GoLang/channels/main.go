/*
// to check links only single time
package main

import (
	"fmt"
	"net/http"
)

func main() { //kind of main routine
	links := []string{"http://golang.org", "http://amazon.com", "http://stackoverflow.com", "http://facebook.com", "http://google.com"}

	fmt.Println(links)
	c := make(chan string) //created channel of string

	for _, url := range links { //passing the slices 1 by 1 in url variable

		go checklink(url, c) //passing channel into child go rotine,so that main routine is aware of the child go routine

	}

	//in order to print every link present in the slice, we used "for loop" otherwise it will only print the faster
	// response link and exit from the program
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c) //blocking call for main go routine means this line will wait for channel to receive msg and then print it on console
	}
} //end of main routine

func checklink(url string, c chan string) {
	_, err := http.Get(url) //blocking call for child go routine (here control will go back to main routine)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- "might be down i think"
		return
	}
	fmt.Println(url, "is up")
	c <- "yes it is up.yayy!"
}

*/

// to check link status recursively with a pause in between

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://golang.org", "http://amazon.com", "http://stackoverflow.com", "http://facebook.com", "http://google.com"}

	c := make(chan string)

	for _, url := range links {

		go checklink(url, c) //blocking call
	}

	// it will iterate over the msgs stored in channel.Similar to:-
	// for {
	//	go checklink(<-c,c)
	//  }
	for l := range c { //channel will recieve msgs and then iterate over it and store in variable l
		go func(url string) { //function literal
			time.Sleep(3 * time.Second) //pause fpr 3 seconds
			checklink(url, c)           //cant use var l here as it is defined at outer scope of this func literal.
		}(l) //call to function literal

		//go checklink(l, c) //blocking call
	}
}

func checklink(url string, c chan string) {
	_, err := http.Get(url) //blocking call
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}
	fmt.Println(url, "is up")
	c <- url
}
