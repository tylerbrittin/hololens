/* Helper functions for MongoDB interaction.
 * Contains a function to return a MongoDB session.
 * Contains a function to submit a "Contact Us" form.
 * Contains all data structs for data submission and retrieval.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package nebulaMongo

import (
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "log"
  "os"
)

// Struct for a detailed Nebula object
type Item struct {
  ID            bson.ObjectId `bson:"_id,omitempty"`
  Category      string
  Item          string
  ItemDesc      string
  Seller        string
  Email         string
  Price         string
  Model         string
  Texture       string
}

// Struct for an Item being Edited
type EItem struct {
  ID            string
  Category      string
  Item          string
  ItemDesc      string
  Seller        string
  Email         string
  Price         string
  Model         string
  Texture       string
}

// Struct for a simple Nebula object
type Simple struct {
  ID            bson.ObjectId `bson:"_id,omitempty"`
  Item          string
  Category      string
  Model         string
  Texture       string
}

// Struct for user info
type UserInfo struct {  
  Username      string
  Password      string
  Firstname     string
  Lastname      string
  Email         string
}

// Struct for adding a user
type UserAddition struct {
  ID            string `bson:"_id,omitempty"`
  Username      string
  Password      string
  Firstname     string
  Lastname      string
  Email         string
}

// Struct for Contact Form
type ContactForm struct {
  Username      string
  Phone         string
  Email         string
  Message       string
}

// Struct for category names
type Cats struct {
  Categories    []string
}

// Struct for adding category
type NewCat struct {
  Name    string
}

// Struct for username check
type Check struct {
  Username  string
  Taken     bool
}

// Struct for deleting an item
type ItemDel struct {
  Seller    string
  Category  string
  ID        string
}

// Get a MongoDB session
func GetSession() *mgo.Session {
  s, err := mgo.Dial("mongodb://localhost")
  if err != nil {
    return nil
  } else {
    return s
  }
}

// Insert new Contact Form
func InsertContact(s *mgo.Session, cf ContactForm) bool {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("contact").C("forms")

  err := c.Insert(cf)

  res := false
  if err != nil {
    log.Println("ERROR: Failed to insert contact form")
    res = true
  }

  return res
}
