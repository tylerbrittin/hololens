package main

import (
  "os"
  "log"
  "net/http"
)

func main() {
  os.Mkdir("/var/nebulaLogs/apiLogs", 0755)
  rotater := NewRotater("/var/nebulaLogs/apiLogs/api.log")
  log.SetOutput(rotater)

  router := NewRouter()
  log.Fatal(http.ListenAndServe(":5073", router))
}
