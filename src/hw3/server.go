package main

import (
    "log"
    "net/http"
    "gopkg.in/mgo.v2"
)

type MongoConnection struct {
    originalSession *mgo.Session
}

func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":1234", router))
}
