package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newdeck() deck {
	card := deck{}

	cardsuits := []string{"spades", "diamonds", "hearts", "cloves"}
	cardvalues := []string{"Ace", "one", "two", "three"}
	for _, s := range cardsuits {
		for _, v := range cardvalues {
			card = append(card, v+" of "+s)
		}
	}

	return card
}

func (d deck) print() {

	for i, c := range d {
		fmt.Println(i, c)
	}
}

func (d deck) tostring() string {
	return (strings.Join([]string(d), ","))
}

func (d deck) savetofile(fname string) error {
	return ioutil.WriteFile(fname, []byte(d.tostring()), 0666)
}

func deckfromfile(fname string) deck {
	bs, err := ioutil.ReadFile(fname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}
