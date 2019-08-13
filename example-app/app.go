package main

import (
	"net/http"
)

func response(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello world`))
}

func main() {
	http.HandleFunc("/", response)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}