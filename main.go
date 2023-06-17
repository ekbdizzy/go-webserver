package main

import "net/http"

type myHandler struct {
}

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func main() {
	err := http.ListenAndServe(":3000", myHandler{})
	if err != nil {
		panic(err)
	}
}
