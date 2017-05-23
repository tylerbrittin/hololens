# Suite of Regression Tests for all endpoints in the
# Nebula Backend System/REST API
#
# Code written by:
# Tim Monfette (tjm354)

from output import colorize, getResults
import requests
import json

# Array of all the test cases
CASES = ["Index", "AddUser", "CheckUsername", "AddItem", "EditItem", "UserItems",
"DeleteItem", "UserInfo", "AddCat", "GetCats", "GetItems", "Details","Contact"]

# Print out final results of Passed/Failed
def finalResults():
  results = getResults() 
  failRatio = results[2] / results[0]
  passRatio = results[1] / results[0]

  print colorize("Number of endpoints that Passed: " + str(results[1]), True)
  print colorize("Number of endpoints that Failed: " + str(results[2]), False)

  # 80% passed is a successful regression test
  print
  if (passRatio < .8):
    print colorize("Too many endpoints Failed, Regression Testing Failed.", False)
    print colorize("Percentage of endpoints that Failed: " + str(failRatio*100) + "%", False)
  else:
    print colorize("At least 80% of endpoints Passed, Regression Testing Passed", True)
    print colorize("Percentage of endpoints that Passed: " + str(passRatio*100) + "%", True)

# Test the / endpoint
def testIndex():
  status = False
  r = requests.get("http://localhost:5073/")

  if ("Welcome! This is the API for the Nebula Shopping portal!" in r.text) and (r.status_code == 200):
    status = True

  print colorize("Testing endpoint /", status)

# Test the /adduser endpoint
def testAddUser():
  status = False
  payload = { "username" : "timmonfette", "password" : "badpass", "firstname" : "Tim", "lastname" : "Monfette", "email" : "monfette.tim@email.com" }
  r = requests.post("http://localhost:5073/adduser", data=json.dumps(payload))

  if (r.status_code == 201):
    status = True

  print colorize("Testing endpoint /adduser", status)

# Test the /checkusername endpoint
def testCheckUsername():
  status = False
  r1 = requests.get("http://localhost:5073/checkusername/timmonfette")
  r2 = requests.get("http://localhost:5073/checkusername/tylerbrittin")

  data1 = json.loads(r1.text)
  data2 = json.loads(r2.text)

  if (data1["Taken"] == True) and (data2["Taken"] == False):
    status = True

  print colorize("Testing endpoint /checkusername", status)

# Test the /additem endpoint
def testAddItem():
  status = False
  payload = { "category" : "electronics", "seller" : "timmonfette", "price" : "$1000", "model" : "stuffy_mod_timmonfette", "texture": "stuffy_tex_timmonfette", "email" : "monfette.tim@email.com", "item" : "brand new stuff", "itemdesc" : "stuffy stuff" }

  r = requests.post("http://localhost:5073/additem", data=json.dumps(payload))
  r2 = requests.get("http://localhost:5073/getitems/electronics")

  data = json.loads(r2.text)

  # Make sure addition was successful
  if (r.status_code == 201 and len(data) == 2):
    status = True

  print colorize("Testing endpoint /additem", status)
  
# Test the /edititem endpoint
def testEditItem():
  status = False
  r = requests.get("http://localhost:5073/getitems/electronics")

  payload = json.loads(r.text)

  if (payload[0]["Item"] == "brand new stuff"):
    ID = payload[0]["ID"]
  else:
    ID = payload[1]["ID"]

  r = requests.get("http://localhost:5073/getdetails/electronics/"+ID)

  payload = json.loads(r.text)
  payload["Price"] = "$5000"

  # Make sure the editing was successful by checking the item price
  r = requests.post("http://localhost:5073/edititem", data=json.dumps(payload))
  r2 = requests.get("http://localhost:5073/getdetails/electronics/"+ID)
  r3 = requests.get("http://localhost:5073/getuseritems/timmonfette")

  data2 = json.loads(r2.text)
  data3 = json.loads(r3.text)

  if (data2["Price"] == "$5000" and data3[0]["Price"] == "$5000"):
    status = True

  print colorize("Testing endpoint /edititem", status)

# Test the /getuseritems endpoint
def testUserItems():
  status = False
  r = requests.get("http://localhost:5073/getuseritems/timmonfette")

  data = json.loads(r.text)

  # Make sure the user has 1 item and it is correct
  if (data[0]["Category"] == "electronics" and data[0]["Item"] == "brand new stuff" and data[0]["ItemDesc"] == "stuffy stuff" and
    data[0]["Seller"] == "timmonfette" and data[0]["Email"] == "monfette.tim@email.com" and data[0]["Price"] == "$5000" and
    data[0]["Model"] == "stuffy_mod_timmonfette" and data[0]["Texture"] == "stuffy_tex_timmonfette" and len(data) == 1):
    status = True

  print colorize("Testing endpoint /getuseritems", status)

# Test the /deleteitem endpoint
def testDeleteItem():
  status = False
  r = requests.get("http://localhost:5073/getuseritems/timmonfette")
  data = json.loads(r.text)

  ID = data[0]["ID"]

  payload = { "seller" : "timmonfette", "category" : "electronics", "id" : ID}

  # Delete the item
  r = requests.post("http://localhost:5073/deleteitem", data=json.dumps(payload))
  r2 = requests.get("http://localhost:5073/getuseritems/timmonfette")
  r3 = requests.get("http://localhost:5073/getitems/electronics")
 
  data3 = json.loads(r3.text)

  # Verify it's deleted on main Nebula side and User item side
  if (r.status_code == 201 and r2.text.strip() == "null"  and len(data3) == 1):
    status = True

  print colorize("Testing endpoint /deleteitem", status)
  
# Test the /getuserinfo endpoint
def testUserInfo():
  status = False
  r = requests.get("http://localhost:5073/getuserinfo/timmonfette")

  data = json.loads(r.text)

  if (data["Username"] == "timmonfette" and data["Password"] == "badpass" and data["Firstname"] == "Tim" and
    data["Lastname"] == "Monfette" and data["Email"] == "monfette.tim@email.com"):
    status = True

  print colorize("Testing endpoint /getuserinfo", status)

# Test the /addcat endpoint
def testAddCat():
  status = False
  payload = { "name" : "testColl" }
  r = requests.post("http://localhost:5073/addcat", data=json.dumps(payload))

  if (r.status_code == 201):
    status = True

  print colorize("Testing endpoint /addcat", status)

# Test the /getcats endpoint
def testGetCats():
  status = False
  r = requests.get("http://localhost:5073/getcats")

  data = json.loads(r.text)
  correct = ["books", "electronics", "furniture", "music", "testColl"]

  if (data["Categories"] == correct):
    status = True 

  print colorize("Testing endpoint /getcats", status)

# Test the /getitems endpoint
def testGetItems():
  status = False
  r = requests.get("http://localhost:5073/getitems/furniture")

  data = json.loads(r.text);
  
  # Test an item from the furniture category
  if (len(data) == 2 and data[0]["Item"] == "New Lamp" and data[0]["Category"] == "furniture" and
    data[0]["Model"] == "lamp_mod_username" and data[0]["Texture"] == "lamp_tex_username"):
    status = True

  print colorize("Testing endpoint /getitems", status)

# Test the /getdetails endpoint
def testDetails():
  status = False
  r = requests.get("http://localhost:5073/getitems/furniture")

  data = json.loads(r.text)
  ID = data[0]["ID"]

  r = requests.get("http://localhost:5073/getdetails/furniture/"+ID)

  data = json.loads(r.text)

  # Make sure the item's details are complete
  if (data["Item"] == "New Lamp" and data["Category"] == "furniture" and data["ItemDesc"] == "new" and
    data["Seller"] == "Mr. Lamp" and data["Email"] == "lamplover@lamp.com" and data["Price"] == "$30" and
    data["Model"] == "lamp_mod_username" and data["Texture"] == "lamp_tex_username"):
    status = True

  print colorize("Testing endpoint /getdetails", status)

# Test the /contactus endpoint
def testContact():
  status = False
  payload = { "username" : "timmonfette", "phone" : "testNum", "email" : "test@test.com", "message" : "Testing contact form" }
  r = requests.post("http://localhost:5073/contactus", data=json.dumps(payload))

  if (r.status_code == 201):
    status = True

  print colorize("Testing endpoint /contactus", status)
