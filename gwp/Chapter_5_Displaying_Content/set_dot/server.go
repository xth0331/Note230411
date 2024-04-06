package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("gwp/Chapter_5_Displaying_Content/set_dot/tmpl.html")
	//t := template.Must(template.ParseFiles("gwp/Chapter_5_Displaying_Content/set_dot/tmpl.html"))
	//t.Execute(w, "Hello")
	//	tmpl := `<!DOCTYPE html>
	//<html>
	//　<head>
	//　　<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	//　　<title>Go Web Programming</title>
	//　</head>
	//　<body>
	//　　{{ . }}
	//　</body>
	//</html>
	//`
	//	t := template.New("tmpl1.html")
	//	t, _ = t.Parse(tmpl)
	//	t.Execute(w, "Hello World!")
	t, _ := template.ParseFiles("gwp/Chapter_5_Displaying_Content/set_dot/tmpl.html")
	t.Execute(w, "hello")
}
func processExec(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("gwp/Chapter_5_Displaying_Content/set_dot/t1.html", "gwp/Chapter_5_Displaying_Content/set_dot/t2.html")
	t.ExecuteTemplate(w, "t1.html", "Hello World!")
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/processexec", processExec)
	server.ListenAndServe()
}
