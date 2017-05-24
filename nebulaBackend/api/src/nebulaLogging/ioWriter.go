/* Handles all writing and maintaining of Nebula REST API Logs.
 * Auto rotates logs (and timestamps old ones) on restart or
 * when the log reaches 10 MB.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaLogging

import (
  "os"
  "sync"
  "time"
  "fmt"
)

type RotateWriter struct {
  lock        sync.Mutex
  filename    string
  fp          *os.File
}

// Return size of file in Megabytes
// Errors opening files are app-ending because it prevents log rotation
func GetFileSize(filepath string) int64 {
  file, err := os.Open(filepath)
  if err != nil {
    fmt.Print("\nERROR: Can not open file to get Size!\n")
    os.Exit(1)
  }

  defer file.Close()

  fi, err := file.Stat()
  if err != nil {
    fmt.Print("\nERROR: Can not get file Size!\n")
    os.Exit(1)
  }

  return fi.Size()
}

// Create a new log rotater
func NewRotater(filename string) *RotateWriter {
  w := &RotateWriter{filename: filename}
  err := w.Rotate()
  if err != nil {
    fmt.Print("\nERROR: Error creating log!\n")
    os.Exit(1)
  }

  return w
}

// Write to the log
func (w *RotateWriter) Write(output []byte) (int, error) {
  w.lock.Lock() 

  size := GetFileSize(w.filename)
  mbSize := size / 1048576
  if mbSize >= 10 {
    w.lock.Unlock()
    w.Rotate()
    w.lock.Lock()
  }

  w.lock.Unlock()
  return w.fp.Write(output)
}

// Rotate a log when it reaches 10 Megabytes
func (w *RotateWriter) Rotate() (err error) {
  w.lock.Lock()
  defer w.lock.Unlock()

  if w.fp != nil {
    err = w.fp.Close()
    w.fp = nil
    if err != nil {
      return
    } 
  }

  _, err = os.Stat(w.filename)
  if err == nil {
    err = os.Rename(w.filename, w.filename+"."+time.Now().Format(time.RFC3339))
    if err != nil {
      return
    }
  }

  w.fp, err = os.Create(w.filename)
  return
}
