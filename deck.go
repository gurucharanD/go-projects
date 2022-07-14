package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck
// which is a slice of strings

type deck []string // this type deck borrows all the properties and behaviour from string

// create a new deck of cards
func newDeck() deck{
	cards := deck{}
	suits := []string{"spades","hearts","diamonds","clubs"}
	values := []string{"ace","two","three","four"}

	for _,s := range suits{
		for _,v:= range values{
			cards = append(cards, v+" of " + s)
		}
	}
	return cards
}


// take existing set of cards, and no of cards to deal with 
// and return a new deck of cards

func deal(d deck, handsize int) (deck,deck){
	return d[:handsize],d[handsize:]
}

// print is a function and a reciever
func (d deck) print(){
	for i,card := range d{
		fmt.Println(i,card)
	}
}

// saveToFile : save a list of cards to file on
// the local machine
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName,[]byte(d.toString()),0666) 
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

func newDeckFromFile(fileName string) deck{
	bs,err := ioutil.ReadFile(fileName)

	if err != nil{
		// log the error and entirely quit the programm
		fmt.Println("error in newDeckFromFile",err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs),",")
	return deck(ss)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for index := range d{ 
		newPos := r.Intn(len(d)-1)
		d[index],d[newPos] = d[newPos],d[index]
	}
	return  
}



