package main

import (
	"fmt"
	"net/http"
)

//	func setCookie(w http.ResponseWriter, r *http.Request) {
//		c1 := http.Cookie{
//			Name:     "first_cookie",
//			Value:    "Go Web Programming",
//			HttpOnly: true,
//		}
//		c2 := http.Cookie{
//			Name:     "second_cookie",
//			Value:    "manning Publications Go",
//			HttpOnly: true,
//		}
//		w.Header().Set("Set-Cookie", c1.String())
//		w.Header().Add("Set-Cookie", c2.String())
//	}
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Manning Publications Go",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

//	func getCookie(w http.ResponseWriter, r *http.Request) {
//		h := r.Header["Cookie"]
//		fmt.Fprintln(w, h)
//	}
func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the Cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)

}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
