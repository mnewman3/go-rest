package main

import (
    "log"
    "net/http"
    "gopkg.in/mgo.v2"
)

const CONNECTIONSTRING = "mongodb://localhost"

type MongoConnection struct {
    originalSession *mgo.Session
}

func main() {

	// create student handler
	StudentHandler := NewStudentHandler(getSession())

	// create our routes using that handler
    routes := CreateRoutes(&StudentHandler)

    // create router using routes created before
    router := NewRouter(routes)

    // start serever on port 1234
    log.Fatal(http.	(":1234", router))
}

// dials our connection string for our db and returns the session
func getSession() *mgo.Session {
	// connect to db
	s, err := mgo.Dial(CONNECTIONSTRING)

	// check for connection error
	if err != nul {
		panic(err)
	}

	return s
}