/* Generic use functions for the Nebula REST API.
 * These functions are route handlers or utility
 * functions that didn't belong in any other
 * category of function.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaHTTPGeneric

import ( 
  "nebulaMongo"
  "encoding/json"
  "io/ioutil"
  "net/http"
  "log"
  "fmt"
  "io" 
)

var contactForm nebulaMongo.ContactForm

// Index Handler
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome! This is the API for the Nebula Shopping portal!\n")
  fmt.Fprint(w, "Current version: 2.3\n")
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

// Handler for Submitting a Contact form
func ContactUs(w http.ResponseWriter, r *http.Request) {

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
  if err := json.Unmarshal(body, &contactForm); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422)
    if err := json.NewEncoder(w).Encode(err); err != nil {
      log.Println("ERROR: JSON gotten from POST request is unprocessable!")
    }
  }

  // Insert Contact Form
  session := nebulaMongo.GetSession()
  res := nebulaMongo.InsertContact(session, contactForm)

  // Return appropriate HTTP code
  if res == true {
    log.Println("ERROR: Failed to insert JSON into Mongo!")
    SendResponse(w, res)
  } else {
    log.Println("INFO: JSON successfully inserted")
    SendResponse(w, res)
  }
}
