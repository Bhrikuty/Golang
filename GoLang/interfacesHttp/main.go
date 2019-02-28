package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	res, err := http.Get("http://google.com") //syntax -> Func get(url string) (res *Response, err Error)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	// fmt.Println(res) // this o/p doesnt print the body of the request

	/*  //Read func here will read the source (ie google.com here) and put the data into byte slice

	bs := make([]byte, 99999)  // this giant number is bcoz Read func is not designed to automatically resize the slice if slice is already full.
	//So if we assign it with 0 element then read func will say that slice is already full

	res.Body.Read(bs) //response struct and then the body property and it has access to read function

	fmt.Println(string(bs))
	*/
	io.Copy(os.Stdout, res.Body) //here "Copy" func is taking data from reader interface and writing to write interface and printing it.
}
