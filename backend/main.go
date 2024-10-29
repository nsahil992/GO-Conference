package main

import (
    "log"
    "net/http"
)

func main() {
    err := initDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    http.HandleFunc("/api/book-ticket", bookTicketHandler)
    http.HandleFunc("/api/tickets", getTicketsHandler)

    log.Println("Server started at http://localhost:9000")
    log.Fatal(http.ListenAndServe(":9000", nil))
}
