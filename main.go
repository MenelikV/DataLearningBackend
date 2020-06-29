package main

import (
    "github.com/gorilla/mux"
    "io"
    "net/http"
)

func IndexHandler(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Hello, Mux!\n")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", IndexHandler)
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}
