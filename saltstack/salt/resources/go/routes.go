package main

import  "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Now",
        "GET",
        "/now",
        Now,
    },
    Route{
        "Check",
        "GET",
        "/check",
        Check,
    },
    Route{
       "NameCreate",
       "POST",
       "/later",
       NameCreate,       
    },
     Route{
       "Invalid",
       "GET",
       "/later",
       Invalid,       
    },   
}
