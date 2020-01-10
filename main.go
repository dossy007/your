package main

import (
	"github.com/dossy007/your/handle"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle.Showdb)

	http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", http.FileServer(http.Dir("stylesheet/"))))
	http.ListenAndServe(":8080", nil)
}
