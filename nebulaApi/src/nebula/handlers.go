package main

import (
  "encoding/json" 
  "net/http"
  "fmt"
  "log"
  "os"
)

var items []Item

// Index Handler
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome! This is the API for the Nebula Shopping portal!\n")
  fmt.Fprint(w, "Current version: 1.0\n")
}

// Handler for adding object to DB - POST request
func AddObject(w http.ResponseWriter, r *http.Request) {

}

// Handler for Books
func GetBooks(w http.ResponseWriter, r *http.Request) { 
  session := GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  items = GetColl(session, "books")

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(items); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  }
}
