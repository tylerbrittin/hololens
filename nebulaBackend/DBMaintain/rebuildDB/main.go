// Go program to rebuild the testing DB.
// This is used to run the Regression suite on a clean DB.
// Assumes you're using a testing version of the API that is
//   configured to use the Testing DBs in MongoDB.
//
// To execute, just move the "rebuildDB" folder to the src/
//   folder of the Go API. Compile the code and then run from bin/
//
// Code written by:
// Tim Monfette (tjm354)

package main

import (
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "fmt"
)

// Struct for a Nebula object
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

func main() {
  DropDB() 
  BuildDB()
}

// Drop the Testing Databases
func DropDB() {
  s := GetSession()
  
  if s == nil {
    fmt.Println("ERROR: Can't get Mongo Session") 
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  err := s.DB("nebulaTest").DropDatabase()
  if err != nil {
    panic(err)
  }

  err = s.DB("contactTest").DropDatabase()
  if err != nil {
    panic(err)
  }

  err = s.DB("usersTest").DropDatabase()
  if err != nil {
    panic(err)
  }
}

// Rebuild the Testing Databases
// Uses some basic, template items for testing
func BuildDB() {
  s := GetSession()

  if s == nil {
    fmt.Println("ERROR: Can't get Mongo Session")
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  // Get a list of template items to insert
  toInsert := ConstructRecords()
  
  // Insert the testing items
  for _, i := range toInsert {
    collName := i.Category
    c := s.DB("nebulaTest").C(collName)
    err := c.Insert(i)

    if err != nil {
      panic(err)
    }
  }
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

// Construct records for the DB
func ConstructRecords() [5]Item {
  var holder [5]Item
  var elec Item
  var book Item
  var musc Item
  var furn1 Item
  var furn2 Item

  elec.Category = "electronics"
  elec.Item = "Example TV"
  elec.ItemDesc = "used"
  elec.Seller = "Mr. Electronic"
  elec.Email = "newmail@electronics.com"
  elec.Price = "$300"
  elec.Model = "tv_mod_username"
  elec.Texture = "tv_tex_username"

  book.Category = "books"
  book.Item = "Example Textbook"
  book.ItemDesc = "new"
  book.Seller = "Mrs. Books"
  book.Email = "bookeep@library.com"
  book.Price = "$150"
  book.Model = "textbook_mod_username"
  book.Texture = "textbook_tex_username"

  musc.Category = "music"
  musc.Item = "Wu Tang Album"
  musc.ItemDesc = "new"
  musc.Seller = "Wu Tang Clan"
  musc.Email = "wutang@clan.com"
  musc.Price = "$1000000"
  musc.Model = "wutang_mod_username"
  musc.Texture = "wutang_tex_username"

  furn1.Category = "furniture"
  furn1.Item = "New Lamp"
  furn1.ItemDesc = "new"
  furn1.Seller = "Mr. Lamp"
  furn1.Email = "lamplover@lamp.com"
  furn1.Price = "$30"
  furn1.Model = "lamp_mod_username"
  furn1.Texture = "lamp_tex_username"

  furn2.Category = "furniture"
  furn2.Item = "Vacuum"
  furn2.ItemDesc = "used"
  furn2.Seller = "Stephen Chiou"
  furn2.Email = "stephen@chiou.com"
  furn2.Price = "$5"
  furn2.Model = "http://www.nebulashop.net/uploads/model_mod_ktest.obj"
  furn2.Texture = "http://www.nebulashop.net/uploads/model_material0000_map_Kd_tex_ktest.png"

  holder[0] = elec
  holder[1] = book
  holder[2] = musc
  holder[3] = furn1
  holder[4] = furn2
  return holder
}
