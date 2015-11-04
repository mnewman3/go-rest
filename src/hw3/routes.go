package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route


func CreateRoutes(h *StudentHandler) Routes {
    return Routes{
        Route{
            "Get Student",
            "GET",
            "/Student",
            h.GetStudent,
        },
        Route{
            "Add Student",
            "POST",
            "/Student",
            h.AddStudent,
        },
        Route{
            "Remove Student",
            "DELETE",
            "/Student",
            h.RemoveStudent,
        },
        Route{
            "Update Student",
            "UPDATE",
            "/Student",
            h.UpdateStudent,
        },
        Route{
            "List Students",
            "LIST",
            "/Student/listall",
            h.ListStudents,
        },



        Route{
            "Beer",
            "GET",
            "/drink",
            Beer,
        },
    }
}
