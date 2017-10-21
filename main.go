package main

import "net/http"

func Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, I'm here!"))
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":80", nil)
}
