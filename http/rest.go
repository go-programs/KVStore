package http

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	"github.com/go-programs/key-value-store/store"
)

func Handler(s store.Service) http.Handler {
	r:= mux.NewRouter()
	r.HandleFunc("/",home(s))
	r.HandleFunc("/v1/key/{key}",putHandler(s)).Methods("PUT")
	r.HandleFunc("/v1/key/{key}",getHandler(s)).Methods("GET")
	r.HandleFunc("/v1/key/{key}",delHandler(s)).Methods("DEL")
	return r
}

func home(s store.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	// Handling 404 error
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Welcome"))
	}
}

func putHandler(s store.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		v, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if(err!=nil) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		s.Put(vars["key"],string(v))
		w.WriteHeader(http.StatusCreated)
	}
}

func getHandler(s store.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		v, err := s.Get(vars["key"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(v))
	}
}

func delHandler(s store.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		s.Del(vars["key"])
		w.WriteHeader(http.StatusOK)
	}
}