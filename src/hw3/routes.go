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
        "/Student",
        GetStudent,
    },
    Route{
        "Add Student",
        "POST",
        "/Student",
        AddStudent,
    },
    Route{
        "Remove Student",
        "DELETE",
        "/Student",
        RemoveStudent,
    },
    Route{
        "Update Student",
        "UPDATE",
        "/Student",
        UpdateStudent,
    },
    Route{
        "List Students",
        "LIST",
        "/Student/listall",
        ListStudents,
    },



    Route{
        "Beer",
        "GET",
        "/drink",
        Beer,
    },
}
