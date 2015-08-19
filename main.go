package main

import "net/http"

func main() {
	println("Running HTTP Server on port: 3000")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":3000", nil)
}
