package main

import (
	"fmt"
)
type Bot interface{
	getGreetings() string
}
type eBot struct{}
type sBot struct{}

func main(){
	eb := eBot{}
	sb := sBot{}
	printGreeting(sb)
	printGreeting(eb)
}

func (eb eBot) getGreetings() string{
	return "Hi"
}

func (sb sBot) getGreetings() string{
	return "Hola"
}

func printGreeting(b Bot){
	fmt.Println(b.getGreetings())
}

// func printGreeting(sb sBot){
// 	fmt.Println(sb.getGreetings())
// }
