package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	rand.New(rand.NewSource(time.Now().Unix()))
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("gwp/Chapter_5_Displaying_Content/nested2/layout.html", "gwp/Chapter_5_Displaying_Content/nested2/red_hello.html")
	} else {
		t, _ = template.ParseFiles("gwp/Chapter_5_Displaying_Content/nested2/layout.html")
		//t, _ = template.ParseFiles("gwp/Chapter_5_Displaying_Content/nested2/layout.html", "gwp/Chapter_5_Displaying_Content/nested2/blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}
func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
