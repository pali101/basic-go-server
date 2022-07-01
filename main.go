package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method - not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started at port 8080.")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
}