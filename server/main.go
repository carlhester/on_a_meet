package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var status bool = false
var updatedAt time.Time

func update(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	key, ok := r.URL.Query()["key"]
	if ok {
		if key[0] == "true" {
			if !status {
				updatedAt = time.Now()
			}
			status = true
			return
		}
	}
	if status {
		updatedAt = time.Now()
	}
	status = false
	fmt.Fprintf(w, "")
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	bgcolor := "red"
	msg := "ON A CALL"
	if !status {
		bgcolor = "green"
		msg = "FREE"
	}

	html := fmt.Sprintf(`
	<!DOCTYPE html><html>
	<head><meta http-equiv="refresh" content="30"></head>
   	<body style="background-color:%s;">
	<p style="font-size: 4rem; color: #FFFFFF; text-align: center;">%s</p>
	<p style="font-size: 1rem; color: #FFFFFF; text-align: center;">Since: %s</p>
    </body></html>`, bgcolor, msg, updatedAt.Format(time.RFC1123))
	fmt.Fprint(w, html)
}

func main() {

	addr := "0.0.0.0:8099"
	updatePath := "/chupdate"

	http.HandleFunc("/", handler)
	http.HandleFunc(updatePath, update)
	log.Printf("Listening on: %s\n", addr)
	log.Printf("UpdatePath: %s\n", updatePath)
	http.ListenAndServe(addr, nil)

}
