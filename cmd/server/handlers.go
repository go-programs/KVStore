
package main

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"github.com/go-programs/key-value-store/cmd/store"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Handling 404 error
	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
	w.Write([]byte("Welcome"))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if(err!=nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	store.Put(vars["key"],string(v))
	w.WriteHeader(http.StatusCreated)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, err := store.Get(vars["key"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(v))
}

func delHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	store.Del(vars["key"])
	w.WriteHeader(http.StatusOK)
}