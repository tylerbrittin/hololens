/* Main program for the Nebula REST API.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package main

import (
  "nebulaLogging"
  "os"
  "log"
  "net/http"
)

func main() {
  os.Mkdir("/var/nebulaLogs/apiLogs", 0755)
  rotater := nebulaLogging.NewRotater("/var/nebulaLogs/apiLogs/api.log")
  log.SetOutput(rotater)

  router := NewRouter()
  log.Fatal(http.ListenAndServe(":5073", router))
}
