package main

import (
	"net/http"
	"log"
	http2 "github.com/go-programs/key-value-store/http"
	kvstore "github.com/go-programs/key-value-store"
	"github.com/go-programs/key-value-store/logger/file"
)


func main()	{
	r := http2.Handler(file.NewFileEventLogger())
	log.Println("Starting server on port 8080")
	err:= http.ListenAndServe(":8080",r)
	if err != nil {
		log.Fatal(err)
	}
}