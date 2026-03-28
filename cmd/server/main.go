package main

import (
	"net/http"
	"resume-maker-pai/internal/service"
	"encoding/json"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.CreateResume())
	fmt.Fprintln(w, "Hello world")
}

func main(){
	http.HandleFunc("/create-resume", handler)

	http.ListenAndServe(":8000", nil)
}