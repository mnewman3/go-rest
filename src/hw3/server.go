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
    routes := CreateRoutes(StudentHandler)

    // create router using routes created before
    router := NewRouter(routes)

    // start serever on port 1234
    log.Fatal(http.ListenAndServe(":1234", router))
}

// dials our db using connection string for our db and returns the session
func getSession() *mgo.Session {
	// connect to db
	s, err := mgo.Dial(CONNECTIONSTRING)

	// check for connection error
	if err != nil {
		panic(err)
	}

	studentCollection := s.DB("StudentInfoDB").C("StudentCollection")
	if studentCollection == nil {
		panic(err)
	}

	// create index on netId so no duplicates in database
	index := mgo.Index{
		Key:      []string{"$text:_id"},
		Unique:   true,
		DropDups: true,
	}
	studentCollection.EnsureIndex(index)

	return s
}