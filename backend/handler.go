package main

import (
    "encoding/json"
    "net/http"
)

func bookTicketHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var ticket Ticket
    err := json.NewDecoder(r.Body).Decode(&ticket)
    if err != nil || ticket.NumTickets > availableTickets {
        http.Error(w, "Invalid input or not enough tickets", http.StatusBadRequest)
        return
    }

    query := "INSERT INTO tickets (buyer_name, num_tickets, email, payment_mode) VALUES (?, ?, ?, ?)"
    _, err = db.Exec(query, ticket.BuyerName, ticket.NumTickets, ticket.Email, ticket.PaymentMode)
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    availableTickets -= ticket.NumTickets // Decrement available tickets
    w.WriteHeader(http.StatusCreated)
}

func getTicketsHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, buyer_name, num_tickets, email, payment_mode FROM tickets")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tickets []Ticket
    for rows.Next() {
        var ticket Ticket
        err := rows.Scan(&ticket.ID, &ticket.BuyerName, &ticket.NumTickets, &ticket.Email, &ticket.PaymentMode)
        if err != nil {
            http.Error(w, "Database error", http.StatusInternalServerError)
            return
        }
        tickets = append(tickets, ticket)
    }

    json.NewEncoder(w).Encode(tickets)
}
