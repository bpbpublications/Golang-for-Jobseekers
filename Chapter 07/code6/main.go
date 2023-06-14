package main

import "net/http"

func specialHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only accepts GET"))
		return
	}
	item := r.URL.Query().Get("item")
	if item == "apple" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	} else if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("expected item to be filled"))
		return
	} else if item != "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("item not found"))
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("the rest of the function is not yet implemented"))
}

func main() {}
