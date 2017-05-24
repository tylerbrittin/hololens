/* Contains all of the functions for doing Item
 * Management for all Nebula Items in MongoDB.
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
    log.Println("ERROR: Failed to delete item from MongoDB")
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
    log.Println("ERROR: Failed to edit item in MongoDB")
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
