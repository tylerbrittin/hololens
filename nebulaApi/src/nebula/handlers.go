package main

import (
  "github.com/gorilla/mux"
  "encoding/json"
  "io/ioutil" 
  "net/http"
  "fmt"
  "log"
  "io"
  "os"
)

//type Respon struct {
//  Result      string
//  Message     string
//}

var simpItems []Simple
var detItem Item
var postItem Item

// Index Handler
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome! This is the API for the Nebula Shopping portal!\n")
  fmt.Fprint(w, "Current version: 1.2\n")
}

// Handler for adding object to DB - POST request
func AddItem(w http.ResponseWriter, r *http.Request) {
  
  // Read Body of POST
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
  if err != nil {
    log.Println("ERROR: Cannot read body from POST request!")
  }
  
  // Close Body
  if err := r.Body.Close(); err != nil {
    log.Println("ERROR: Cannot close body of the POST request!")
  }
  
  // Process JSON
  if err := json.Unmarshal(body, &postItem); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Println("ERROR: JSON gotten from POST request is unprocessable!")
    }
  }

  // Insert item in Mongo
  session := GetSession()
  res := InsertItem(session, postItem)

  w.Header().Set("Access-Control-Allow-Origin", "http://nebulashop.net")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

  // Return appropriate HTTP code
  if res == true {
    log.Println("ERROR: Failed to insert JSON into Mongo!")
    SendResponse(w, res)
  } else {
    log.Println("INFO: JSON successfully inserted")
    //myres := &Respon{
      //Result: "Success",
      //Message: "Successfuly inserted JSON"}
    SendResponse(w, res)
  }
}

// Handler for matching items
func GetItems(w http.ResponseWriter, r *http.Request) { 
  session := GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  vars := mux.Vars(r)
  collName := vars["collName"]

  simpItems = GetColl(session, collName)
 
  w.Header().Set("Access-Control-Allow-Origin", "http://nebulashop.net")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(simpItems); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  }
}

// Handler for item details
func GetDetails(w http.ResponseWriter, r *http.Request) {
  session := GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  vars := mux.Vars(r)
  collName := vars["collName"]
  id := vars["id"]

  detItem = GetItemDets(session, collName, id)

  w.Header().Set("Access-Control-Allow-Origin", "http://nebulashop.net")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(detItem); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  }
}

// Helper Function for sending response after POST
func SendResponse(w http.ResponseWriter, res bool) {
  if res == true {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest) 
  } else {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
  }
}
