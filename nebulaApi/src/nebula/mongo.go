package main

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
}

// Struct for a simple Nebula object
type Simple struct {
  ID            bson.ObjectId `bson:"_id,omitempty"`
  Item          string
  Model         string
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

// Get an entire collection
func GetColl(s *mgo.Session, collName string) []Simple {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("nebula").C(collName)

  var results []Simple
  err := c.Find(nil).All(&results)

  if err != nil {
    log.Printf("FATAL: Can not access "+collName+" collection to get items! Application Closing!")
    os.Exit(1)
  }

  return results
}

// Get individual item details
func GetItemDets(s *mgo.Session, collName string, id string) Item {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("nebula").C(collName)

  var result Item
  err := c.FindId(bson.ObjectIdHex(id)).One(&result)

  if err != nil {
    log.Printf("FATAL: Can not access "+collName+" collection to get item! Application Closing!")
    os.Exit(1)
  }

  return result
}

// Insert an item into MongoDB
func InsertItem(s *mgo.Session, i Item) bool {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  collName := i.Category
  defer s.Close()

  s.SetMode(mgo.Monotonic, true)

  c := s.DB("nebula").C(collName)

  err := c.Insert(i)

  res := false
  if err != nil {
    res = true
  }

  return res
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
