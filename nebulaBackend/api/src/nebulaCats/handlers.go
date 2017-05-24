/* All of the route handlers for endpoints that
 * are related to Categories in Nebula.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaCats

import (
  "nebulaHTTPGeneric"
  "nebulaMongo" 
  "encoding/json"
  "io/ioutil"
  "net/http"
  "log"
  "io"
  "os"
)

// Data structs
var cats nebulaMongo.Cats
var newCat nebulaMongo.NewCat

// Handler for adding a new category of item
func AddCat(w http.ResponseWriter, r *http.Request) {

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
  if err := json.Unmarshal(body, &newCat); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Println("ERROR: JSON gotten from POST request is unprocessable!")
    }
  }

  session := nebulaMongo.GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  collName := newCat.Name
  res := nebulaMongo.CreateCategory(session, collName)

  // Return appropriate HTTP code
  if res == true {
    log.Println("ERROR: Failed to insert JSON into Mongo!")
    nebulaHTTPGeneric.SendResponse(w, res)
  } else {
    log.Println("INFO: JSON successfully inserted")
    nebulaHTTPGeneric.SendResponse(w, res)
  }
}

// Handler for getting categories
func GetCats(w http.ResponseWriter, r *http.Request) {
  session := nebulaMongo.GetSession()

  if session == nil {
    log.Println("ERROR: FATAL: Unable to get MongoDB Connection. Exiting!")
    os.Exit(1)
  }

  cats.Categories = nebulaMongo.GetCategories(session)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(cats); err != nil {
    log.Println("ERROR: Unable to properly encode items from MongoDB as JSON!")
  }
}
