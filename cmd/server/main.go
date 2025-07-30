package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    // API endpoint example
    r.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"message":"Hello from Go backend"}`))
    })

    // Serve static files for the frontend
    fs := http.FileServer(http.Dir("./frontend/dist"))
    r.PathPrefix("/").Handler(fs)

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
