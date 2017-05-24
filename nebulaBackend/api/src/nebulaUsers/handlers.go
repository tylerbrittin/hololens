/* All of the route handlers for endpoints related to
 * User Management of Nebula users.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaUsers

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
var userIts []nebulaMongo.Item
var users nebulaMongo.Cats
var user nebulaMongo.UserInfo
var postUser nebulaMongo.UserInfo

// Handler for getting user info
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
  session := nebulaMongo.GetSession()
  
  if session == nil {
    log.Println("FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  } 
  
  vars := mux.Vars(r)
  username := vars["username"]
  
  user = nebulaMongo.GetUser(session, username)
  
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(user); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  } 
} 

// Handler for getting all items for sale by a user
func GetUserItems(w http.ResponseWriter, r *http.Request) {
  session := nebulaMongo.GetSession()
  
  if session == nil {
    log.Println("FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  } 
  
  vars := mux.Vars(r)
  username := vars["username"]
  
  userIts = nebulaMongo.UserItems(session, username)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(userIts); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  } 
}

// Handler for checking if a username is taken already
func CheckUsername(w http.ResponseWriter, r *http.Request) {
  session := nebulaMongo.GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  vars := mux.Vars(r)
  username := vars["username"]

  users.Categories = nebulaMongo.GetUserList(session)

  var userResult nebulaMongo.Check
  taken := false

  for _, name := range users.Categories {
    if name == username {
      taken = true
    }
  }

  userResult.Username = username
  userResult.Taken = taken

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(userResult); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  }
}

// Handler for adding new user to MongoDB - POST request
func AddUser(w http.ResponseWriter, r *http.Request) {

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
  if err := json.Unmarshal(body, &postUser); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Println("ERROR: JSON gotten from POST request is unprocessable!")
    }
  }

  // Insert user in Mongo
  session := nebulaMongo.GetSession()
  res := nebulaMongo.InsertUser(session, postUser)

  // Return appropriate HTTP code
  if res == true {
    log.Println("ERROR: Failed to insert JSON into Mongo!")
    nebulaHTTPGeneric.SendResponse(w, res)
  } else {
    log.Println("INFO: JSON successfully inserted")
    nebulaHTTPGeneric.SendResponse(w, res)
  }
}
