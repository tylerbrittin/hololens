/* Contains the functions for handling Categories
 * of Nebula Items in MongoDB.
 *
 * Categories of Nebula items are collections inside
 * the Nebula DB.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaMongo

import (
  "gopkg.in/mgo.v2"
  "log"
  "os"
)


// Get a list of all categories of items for sale
func GetCategories(s *mgo.Session) []string {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  cats, err := s.DB("nebula").CollectionNames()

  if err != nil {
    log.Printf("ERROR: Error accessing list of categories!")
  }

  return cats
}

// Create a new category of items to sell
func CreateCategory(s *mgo.Session, name string) bool {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("nebula").C(name)
  err := c.Create(&mgo.CollectionInfo{})

  res := false
  if err != nil {
    log.Printf("ERROR: Can not create "+name+" as a category!")
    res = true
  }

  return res
}
