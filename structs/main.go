package main

import "fmt"
type contacInfo struct{
	email string
	zip int
}

type person struct{
	firstname string
	lastname string
	contact contacInfo
}





func main(){
	person1 := person{"guru","charan",contacInfo{"g@gmail",123}}
	person2 := person{firstname:"guru",lastname:"charan", contact: contacInfo{
		email: "g@g",
		zip: 123,
		},
	}
	var person3 person
	person3.firstname = "new name"
	fmt.Println(person3)
	fmt.Println(person1)
	fmt.Printf("%+v",person2)
	person1.printName()
	person1.updateName("temp")
	person1.printName()

	// person1Pointer := &person1
	// person1Pointer.updateName("temp")
	// person1Pointer.printName()

	colors := map[string] string{
		"red":"#223343",
		"green":"#4335",
	}
	fmt.Println(colors)
	fmt.Println(colors["red"])

	for key,value := range colors{
		fmt.Println(key,value)
	}
}

func (p person) printName(){
	fmt.Println(p.firstname+p.lastname)
	return
}

func (p *person) updateName(name string){
	p.firstname = name
}