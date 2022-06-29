package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func main()	{
	r:= mux.NewRouter()
	r.HandleFunc("/",home)
	r.HandleFunc("/v1/key/{key}",putHandler).Methods("PUT")
	r.HandleFunc("/v1/key/{key}",getHandler).Methods("GET")
	r.HandleFunc("/v1/key/{key}",delHandler).Methods("DEL")
	log.Println("Starting server on port 8080")
	err:= http.ListenAndServe(":8080",r)
	if err != nil {
		log.Fatal(err)
	}
}