package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "This is service 1 project")
    })
    http.ListenAndServe(":8080", nil)
}
