
package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Handling 404 error
	if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
	w.Write([]byte("Welcome"))
}
