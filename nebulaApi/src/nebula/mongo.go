package main

import (
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "log"
  "os"
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

// Get an entire collection
func GetColl(s *mgo.Session, collName string) []Item {
  if s == nil {
    log.Printf("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("nebula").C(collName)

  var results []Item
  err := c.Find(nil).All(&results)

  if err != nil {
    log.Printf("FATAL: Can not access "+collName+" collection to get items! Application Closing!")
    os.Exit(1)
  }

  return results
}

// Insert an item into MongoDB
func InsertItem(s *mgo.Session, i Item) bool {
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
