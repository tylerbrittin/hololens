/* Creates all of the routes/endpoints serviced by
 * the Nebula REST API.
 *
 * Code Written by:
 * Tim Monfette (tjm354)
*/

package main

import (
  "net/http"
  "nebulaHTTPGeneric"
  "nebulaCats"
  "nebulaUsers"
  "nebulaItems"
)

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    nebulaHTTPGeneric.Index,
  },
  Route{
    "AddObject",
    "POST",
    "/additem",
    nebulaItems.AddItem,
  },
  Route{
    "AddCat",
    "POST",
    "/addcat",
    nebulaCats.AddCat,
  },
  Route{
    "GetItems",
    "GET",
    "/getitems/{collName}",
    nebulaItems.GetItems,
  },
  Route{
    "DeleteItem",
    "POST",
    "/deleteitem",
    nebulaItems.DeleteItem,
  },
  Route{
    "EditItem",
    "POST",
    "/edititem",
    nebulaItems.EditItem,
  },
  Route{
    "GetCats",
    "GET",
    "/getcats",
    nebulaCats.GetCats,
  },
  Route{
    "GetDetails",
    "GET",
    "/getdetails/{collName}/{id}",
    nebulaItems.GetDetails,
  },
  Route{
    "GetUserInfo",
    "GET",
    "/getuserinfo/{username}",
    nebulaUsers.GetUserInfo,
  },
  Route{
    "GetUserItems",
    "GET",
    "/getuseritems/{username}",
    nebulaUsers.GetUserItems,
  },
  Route{
    "CheckUsername",
    "GET",
    "/checkusername/{username}",
    nebulaUsers.CheckUsername,
  },
  Route{
    "AddUser",
    "POST",
    "/adduser",
    nebulaUsers.AddUser,
  },
  Route{
    "ContactUs",
    "POST",
    "/contactus",
    nebulaHTTPGeneric.ContactUs,
  },
}
