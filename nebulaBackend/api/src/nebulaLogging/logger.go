/* Creates the logger and logging structure for all
 * logged output from the Nebula REST API.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaLogging

import (
  "log"
  "net/http"
  "time"
)

func Logger(inner http.Handler, name string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()

    inner.ServeHTTP(w, r)

    log.Printf(
      "%s\t%s\t%s\t%s",
      r.Method,
      r.RequestURI,
      name,
      time.Since(start),
    )
  })
}
