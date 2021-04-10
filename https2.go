package main

import (
	"net/http"
	"io"
	"fmt"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS,PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Lenght, Authorization")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return 
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Lenght, Authorization")
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful post")
	}
	if req.Method =="PUT" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return 
		}
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Lenght, Authorization")
		fmt.Printf("%s\n", data)
		io.WriteString(w, "successful put")
	}
	
}


func main() {
	http.HandleFunc("/", Handler)
	
	err := http.ListenAndServe(":8080", nil)
	panic(err)
}
