// main.go

package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Ana HTTP yönlendirici oluşturma
    http.HandleFunc("/", handler)

    // Sunucuyu başlatma
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

// handler fonksiyonu, gelen HTTP isteklerini işler
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}