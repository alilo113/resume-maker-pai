package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/create-resume", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I dont know what the fuck am doing!")
	})
}