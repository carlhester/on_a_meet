package main

import (
	"fmt"
	"net/http"
)

var STATUS bool = false

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	key, ok := r.URL.Query()["key"]
	if ok {
		if key[0] == "true" {
			STATUS = true
			return
		}
	}
	STATUS = false
	fmt.Fprintf(w, "")
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	bgcolor := "red"
	msg := "ON A CALL"
	if !STATUS {
		bgcolor = "green"
		msg = "FREE"
	}

	html := fmt.Sprintf(`
	<!DOCTYPE html><html>
	<head><meta http-equiv="refresh" content="5"></head>
   	<body style="background-color:%s;">
	<p style="font-size: 4rem; color: #FFFFFF; text-align: center;">%s</p>
    </body></html>`, bgcolor, msg)
	fmt.Fprint(w, html)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/chupdate", update)
	http.ListenAndServe(":8099", nil)

}
