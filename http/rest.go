package http

import (
	"net/http"
	"io"
	"github.com/gorilla/mux"
	kvstore "github.com/go-programs/key-value-store"
)

var (
	el kvstore.EventLogger
)

func Handler(eventLogger kvstore.EventLogger) http.Handler {
	el = eventLogger
	r:= mux.NewRouter()
	r.HandleFunc("/",home)
	r.HandleFunc("/v1/key/{key}",putHandler).Methods("PUT")
	r.HandleFunc("/v1/key/{key}",getHandler).Methods("GET")
	r.HandleFunc("/v1/key/{key}",delHandler).Methods("DEL")
	return r
}

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
	kvstore.Put(vars["key"],string(v))
	el.writePut(vars["key"],string(v))
	w.WriteHeader(http.StatusCreated)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, err := kvstore.Get(vars["key"])
	el.writeGet(vars["key"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(v))
}

func delHandler(w http.ResponseWriter, r *http.Request) {
 	vars := mux.Vars(r)
 	kvstore.Del(vars["key"])
 	el.writeDelete(vars["key"])
 	w.WriteHeader(http.StatusOK)
 }
