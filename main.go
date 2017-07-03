package main

import "net/http"

func main() {
	http.HandleFunc("/ps", HandleHTTP)
	http.ListenAndServe(":8080", nil)
}
