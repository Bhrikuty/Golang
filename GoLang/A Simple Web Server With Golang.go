package main

import (
	"log"
	"net/http"
	"strings"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

//the above function also works similar as below function (below func is easy to understand)

func hello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hi there, I love " + message
	w.Write([]byte(message))
}

func main() {
	//call to http.HandleFunc, which tells the http package to handle all requests to the web root ("/") with handler called "hello"

	http.HandleFunc("/", hello)

	//calls http.ListenAndServe, specifying that it should listen on port 8080 on any interface

	log.Fatal(http.ListenAndServe(":8080", nil))

	//ListenAndServe always returns an error
	//In order to log that error we wrap the function call with log.Fatal.
}
