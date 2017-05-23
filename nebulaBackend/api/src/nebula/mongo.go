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
    log.Printf("ERROR: Can not access "+collName+" collection to get items!") 
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
    log.Printf("ERROR: Can not access "+collName+" collection to get item!")
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

  mID := bson.NewObjectId()
  i.ID = mID 

  err := c.Insert(i)

  res := false
  if err != nil {
    res = true
    return res
  }

  c = s.DB("users").C(i.Seller)
  err = c.Insert(i)

  res = false
  if err != nil {
    res = true
  }

  return res
}
// Delete an Item from MongoDB
func DeleteRecord(s *mgo.Session, del ItemDel) bool {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!");
    os.Exit(1)
  }

  collName := del.Category
  username := del.Seller
  mID := bson.ObjectIdHex(del.ID)

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("nebula").C(collName)
  err := c.Remove(bson.M{"_id": mID})
  
  res := false
  if err != nil {
    res = true
    return res
  }

  c = s.DB("users").C(username)
  err = c.Remove(bson.M{"_id": mID})

  res = false
  if err != nil {
    res = true
    return res
  }

  return res
}

// Edit an Item in MongoDB
func EditRecord(s *mgo.Session, ei EItem) bool {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!");
    os.Exit(1)
  }

  collName := ei.Category
  username := ei.Seller
  mID := bson.ObjectIdHex(ei.ID)
  querier := bson.M{"_id": mID}
  change := bson.M{"$set": ei}

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("nebula").C(collName)
  err := c.Update(querier, change)
  
  res := false
  if err != nil {
    res = true
    return res
  }

  c = s.DB("users").C(username)
  err = c.Update(querier, change)

  if err != nil {
    res = true
    return res
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

// Function for getting a users info
func GetUser(s *mgo.Session, name string) UserInfo {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("users").C(name)

  var result UserInfo
  err := c.Find(bson.M{"_id": "userInfo"}).One(&result)
  
  if err != nil {
    log.Printf("ERROR: Can not access user info for "+name+"!")
  }

  return result
}

// Function to get all items for sale by a user
func UserItems(s *mgo.Session, name string) []Item {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  c := s.DB("users").C(name)

  var results []Item
  err := c.Find(bson.M{"_id": bson.M{"$not": bson.RegEx{`userInfo`, ""}}}).All(&results) 

  if err != nil {
    log.Printf("ERROR: Can not access items for sale by "+name+"!")
  }

  return results
}

// Get a list of all current users in the system
func GetUserList(s *mgo.Session) []string {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
  }

  defer s.Close()
  s.SetMode(mgo.Monotonic, true)

  users, err := s.DB("users").CollectionNames()

  if err != nil {
    log.Printf("ERROR: Error accessing list of users!")
  }

  return users
}

// Insert a new user into MongoDB
func InsertUser(s *mgo.Session, u UserInfo) bool {
  if s == nil {
    log.Println("FATAL: Can not access MongoDB! Application Closing!")
    os.Exit(1)
  }

  collName := u.Username
  defer s.Close()

  s.SetMode(mgo.Monotonic, true)

  c := s.DB("users").C(collName)

  var ua UserAddition
  ua.ID = "userInfo"
  ua.Username = collName
  ua.Password = u.Password
  ua.Firstname = u.Firstname
  ua.Lastname = u.Lastname
  ua.Email = u.Email

  err := c.Insert(ua)

  res := false
  if err != nil {
    res = true
  }

  return res
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
    res = true
  }

  return res
}
