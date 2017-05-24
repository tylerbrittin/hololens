/* All of the route handlers for endpoints related to
 * Item management in Nebula.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaItems

import (
  "nebulaHTTPGeneric"
  "nebulaMongo"
  "github.com/gorilla/mux"
  "encoding/json"
  "io/ioutil" 
  "net/http"
  "log"
  "io"
  "os"
)

// Data Structs
var simpItems []nebulaMongo.Simple
var detItem nebulaMongo.Item
var postItem nebulaMongo.Item
var toDelete nebulaMongo.ItemDel
var toEdit nebulaMongo.EItem

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
  session := nebulaMongo.GetSession()
  res := nebulaMongo.InsertItem(session, postItem)

  // Return appropriate HTTP code
  if res == true {
    log.Println("ERROR: Failed to insert JSON into Mongo!")
    nebulaHTTPGeneric.SendResponse(w, res)
  } else {
    log.Println("INFO: JSON successfully inserted") 
    nebulaHTTPGeneric.SendResponse(w, res)
  }
}

// Handler for deleting an Item
func DeleteItem(w http.ResponseWriter, r *http.Request) {

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
  if err := json.Unmarshal(body, &toDelete); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Println("ERROR: JSON gotten from POST request is unprocessable!")
    }
  }

  // Delete Item from MongoDB
  session := nebulaMongo.GetSession()
  res := nebulaMongo.DeleteRecord(session, toDelete)
 
  // Return appropriate HTTP code
  if res == true {
    log.Println("ERROR: Failed to delete Item from Mongo!")
    nebulaHTTPGeneric.SendResponse(w, res)
  } else {
    log.Println("INFO: Item successfully deleted")
    nebulaHTTPGeneric.SendResponse(w, res)
  }
}

// Handler for editing an Item
func EditItem(w http.ResponseWriter, r *http.Request) {
  
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
  if err := json.Unmarshal(body, &toEdit); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Println("ERROR: JSON gotten from POST request is unprocessable!")
    }
  }

  // Edit Item in MongoDB
  session := nebulaMongo.GetSession()
  res := nebulaMongo.EditRecord(session, toEdit)

  // Return appropriate HTTP code
  if res == true {
    log.Println("ERROR: Failed to delete Item from Mongo!")
    nebulaHTTPGeneric.SendResponse(w, res)
  } else {
    log.Println("INFO: Item successfully deleted")
    nebulaHTTPGeneric.SendResponse(w, res)
  }
}

// Handler for matching items
func GetItems(w http.ResponseWriter, r *http.Request) { 
  session := nebulaMongo.GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  vars := mux.Vars(r)
  collName := vars["collName"]

  simpItems = nebulaMongo.GetColl(session, collName)

  // Return all items in that collection 
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(simpItems); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  }
}

// Handler for item details
func GetDetails(w http.ResponseWriter, r *http.Request) {
  session := nebulaMongo.GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  // Get collection name and item ID
  vars := mux.Vars(r)
  collName := vars["collName"]
  id := vars["id"]

  detItem = nebulaMongo.GetItemDets(session, collName, id)

  // Return the data
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(detItem); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  }
}
