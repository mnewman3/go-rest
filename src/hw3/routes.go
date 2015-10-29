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
        "Get Student",
        "GET",
        "/Student/{name}",
        Index,
    },
    Route{
        "Add Student",
        "POST",
        "/Student",
        TodoCreate,
    },
    Route{
        "Remove Student",
        "DELETE",
        "/Student",
        TodoIndex,
    },
    Route{
        "Update Student",
        "UPDATE",
        "/Student",
        TodoCreate,
    },
    Route{
        "List Students",
        "LIST",
        "/Student/listall",
        TodoCreate,
    },
    Route{
        "Beer",
        "GET",
        "/drink",
        Beer,
    },
}
