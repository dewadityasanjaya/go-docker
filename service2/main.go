package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "This is service 2 project")
    })
    http.ListenAndServe(":8081", nil)
}
