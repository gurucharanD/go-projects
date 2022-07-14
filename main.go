package main

import "fmt"

// import "fmt"

func main(){
	cards := newDeck()
	cards.saveToFile("cards")

	// fmt.Println(cards.toString())
	cards.print()

	// d1,d2 := deal(cards,3)
	// d1.print()
	// d2.print()
	
	// newDeckFromFile("cards222").print()
	fmt.Println("____")
	cards.shuffle()
	cards.print()

}

func newCard() string{
	return "new card"
}
 