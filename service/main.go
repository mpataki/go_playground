package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", helloHandler) // Route to handle requests to "/"

    port := "8080"
    fmt.Println("Server is listening on port", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

