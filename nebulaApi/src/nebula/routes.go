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
    "GetItems",
    "GET",
    "/getitems/{collName}",
    GetItems,
  },
  Route{
    "GetDetails",
    "GET",
    "/getdetails/{collName}/{id}",
    GetDetails,
  },
}
