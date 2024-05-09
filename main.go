package main

import (
	"log"
	"net/http"
)

const (
    PORT    = ":8000"
)

func main() {
    router := http.NewServeMux()

    router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, World\n"))
    })

    log.Printf("Server running on http://localhost%v", PORT)
    log.Fatal(http.ListenAndServe(PORT, router))
}
