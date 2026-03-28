package main

import (
	"net/http"
	"resume-maker-pai/internal/service"
	"encoding/json"
)

func main(){
	http.HandleFunc("/create-resume", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(service.CreateResume())
	})

	http.ListenAndServe(":8000", nil)
}