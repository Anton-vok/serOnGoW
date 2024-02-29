package main

import "net/http"

func main() {
	http.Handle("/qr/", http.StripPrefix("/qr/", http.FileServer(http.Dir("files"))))
	http.ListenAndServe(":8080", nil)
}
