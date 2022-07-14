package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"404",http.StatusNotFound)
		return
	}
}

func formHandler(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"ParseForm() err %v",err)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Println(name,address)
	fmt.Fprintf(w,name)
	fmt.Fprintf(w,address)
	
}	

func main(){

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Println("server running on port 8080")
	if err:= http.ListenAndServe(":8080",nil); err!=nil{
		log.Fatal(err)
	}

}

