package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("gwp/Chapter_5_Displaying_Content/random_number/tmpl.html")
	//rand.Seed(time.Now().Unix()) 废弃
	rand.New(rand.NewSource(time.Now().Unix()))
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
