package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

//func index(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Home"))
//}
//func hello(w http.ResponseWriter, r *http.Request) {
//	name := strings.Split(r.URL.Path, "/")[2]
//	w.Write([]byte(fmt.Sprintf("Hello, %s", name)))
//}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index page"))
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
