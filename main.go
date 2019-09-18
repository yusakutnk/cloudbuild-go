package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200)
        fmt.Fprintf(w, "Welcome to Secret Page!!!")
    })

    port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
    }

    http.ListenAndServe(":" + port, nil)
}

