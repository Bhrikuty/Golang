package main

import (
	"fmt"
	"io/ioutil"
)

type page struct {
	title string
	body  []byte
}

func main() {
	p := &page{title: "hellofile", body: []byte("this is a first file")}
	p.save()                     //call to create a file and write into it.
	p1, _ := load("hellofile")   //attempt to read a file whose response gets stored in var p1
	fmt.Println(string(p1.body)) // print the body part of the response on the terminal
}

func (p *page) save() error { //return type is error as writefile returns error
	fname := p.title + ".txt"                    //create a file name of your choice
	return ioutil.WriteFile(fname, p.body, 0600) //write in file the text which was mentioned in the body variable
}

func load(t string) (*page, error) {
	fname := t + ".txt"
	b, err := ioutil.ReadFile(fname) //read file contents and pass into var b
	if err != nil {
		return nil, err
	}
	return &page{title: t, body: b}, nil // then return the var b in a struct
}
