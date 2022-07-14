package main

import (
	"fmt"
	"net/http"
)

func main(){
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}
	c := make(chan string)

	for _,url := range urls{
		// fmt.Println(url)
		// data:= ch		
		go checkLink(url,c)

	}
	
	for i:=0; i<len(urls);i++{
		fmt.Println(<-c)
	}

}

func checkLink(url string,c chan(string)){
	data,err := http.Get(url)
	if err != nil{
		fmt.Println("error with ", err,data)
		c <- url
		return
	}
	fmt.Println(url)
	c <- url
}