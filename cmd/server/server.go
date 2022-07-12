package main

import (
	"net/http"
	"log"
	http2 "github.com/go-programs/key-value-store/http"
)

func main()	{

	r := http2.Handler()
	log.Println("Starting server on port 8080")
	err:= http.ListenAndServe(":8080",r)
	if err != nil {
		log.Fatal(err)
	}
}