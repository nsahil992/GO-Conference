package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
const maxTickets = 100 // Set the maximum tickets available
var availableTickets = maxTickets

func initDB() error {
    var err error
    db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/conference_db")
    if err != nil {
        return err
    }
    return db.Ping()
}

type Ticket struct {
    ID          int
    BuyerName   string
    NumTickets  int
    Email       string
    PaymentMode string
}

