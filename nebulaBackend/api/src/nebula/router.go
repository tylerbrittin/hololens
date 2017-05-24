/* Creates the Mux Router to handle all incoming HTTP
 * requests and route them to the proper handlers.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package main

import (
  "nebulaLogging"
  "net/http"
  "github.com/gorilla/mux"
)

// Create the router
func NewRouter() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)
  for _, route := range routes {
    var handler http.Handler

    handler = route.HandlerFunc
    handler = nebulaLogging.Logger(handler, route.Name)

    router.
      Methods(route.Method).
      Path(route.Pattern).
      Name(route.Name).
      Handler(handler)
  }

  return router
}
