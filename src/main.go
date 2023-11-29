package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func serveFiles(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("URL path: %s\n", r.URL.Path)
    http.ServeFile(w, r, "./src/static/index.html")
}

func main() {
    fmt.Println("Now listening on 8080")
    http.HandleFunc("/", serveFiles)
    err := http.ListenAndServe(":8080", nil)

    if errors.Is(err, http.ErrServerClosed) {
        fmt.Println("Server closed") 
    } else if err != nil {
        fmt.Printf("Error starting server: %s\n", err)
        os.Exit(1)
    }
}
