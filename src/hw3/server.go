package main

import (
    "log"
    "net/http"
)

type MongoConnection struct {
    originalSession *mgo.Session
}

func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":1234", router))
}
