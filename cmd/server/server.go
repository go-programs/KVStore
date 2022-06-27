package main

import (
	"net/http"
	"log"
)

func main()	{
	mux:= http.NewServeMux()
	mux.HandleFunc("/",home)
	log.Println("Starting server on port 8080")
	err:= http.ListenAndServe(":8080",mux)
	if err != nil {
		log.Fatal(err)
	}
}