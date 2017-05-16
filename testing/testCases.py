from output import colorize
import requests
import json

# Array of all the test cases
CASES = ["Index", "AddUser", "CheckUsername", "UserItems", "UserInfo"]

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
    payload = { "username" : "tylerbrittin", "password" : "badpass", "firstname" : "Tyler", "lastname" : "Brittin", "email" : "brittin.tyler@email.com" }
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

    if (data1["Taken"] == False) and (data2["Taken"] == True):
        status = True

    print colorize("Testing endpoint /checkusername", status)

# Test the /getuseritems endpoint
def testUserItems():
    status = False
    r = requests.get("http://localhost:5073/getuseritems/tylerbrittin")

    if (r.text.strip() == "null"):
        status = True

    print colorize("Testing endpoint /getuseritems", status)

# Test the /getuserinfo endpoint
def testUserInfo():
    status = False
    r = requests.get("http://localhost:5073/getuserinfo/tylerbrittin")

    data = json.loads(r.text)

    if (data["Username"] == "tylerbrittin") and (data["Password"] == "badpass") and (data["Firstname"] == "Tyler") and (data["Lastname"] == "Brittin") and (data["Email"] == "brittin.tyler@email.com"):
        status = True

    print colorize("Testing endpoint /getuserinfo", status)

# Test the /addcat endpoint

# Test the /getcats endpoint

# Test the /additem endpoint

# Test the /getitems endpoint

# Test the /getdetails endpoint

# Test the /getuseritems endpoint
