package main

import (
	"encoding/json"
	"fmt"
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

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func postUser(w http.ResponseWriter, r *http.Request) {
	var user User
	//err := json.Unmarshal([]byte(r.Body), &user)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	bytes, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("500 error" + err.Error()))
		return
	}
	w.Write(bytes)

}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of products"))
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	var id string
	id = chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("Id is %s", id)))
	return

}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index page"))
	})

	r.Route("/user/", func(r chi.Router) {
		r.With().Post("/add/", postUser)
	})

	r.Route("/products", func(r chi.Router) {
		r.With().Get("/", getProducts)
		r.With().Get("/{id:^[0-9]+}", getProduct)
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
}
