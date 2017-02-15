package main

import "net/http"

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
    Index,
  },
  Route{
    "AddObject",
    "POST",
    "/additem",
    AddItem,
  },
  Route{
    "AddCat",
    "POST",
    "/addcat",
    AddCat,
  },
  Route{
    "GetItems",
    "GET",
    "/getitems/{collName}",
    GetItems,
  },
  Route{
    "GetCats",
    "GET",
    "/getcats",
    GetCats,
  },
  Route{
    "GetDetails",
    "GET",
    "/getdetails/{collName}/{id}",
    GetDetails,
  },
  Route{
    "GetUserInfo",
    "GET",
    "/getuserinfo/{username}",
    GetUserInfo,
  },
  Route{
    "GetUserItems",
    "GET",
    "/getuseritems/{username}",
    GetUserItems,
  },
}
