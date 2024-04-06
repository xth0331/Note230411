package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("gwp/Chapter_5_Displaying_Content/include/t1.html", "gwp/Chapter_5_Displaying_Content/include/t2.html")
	t.Execute(w, "Hello World!")
}
func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
