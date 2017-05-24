/* Contains all functions for managing Nebula Users in
 * MongoDB.
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
    log.Println("ERROR: Failed to create new user")
    res = true
  }

  return res
}
