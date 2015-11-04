package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    // "io"
    // "io/ioutil"

    "hw3/models"
    "github.com/gorilla/mux"
)

/*
    // func Index(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintln(w, "Welcome!")
    // }

    // func TodoIndex(w http.ResponseWriter, r *http.Request) {
    //     // todos := Todos{
    //     //     Todo{Name: "Write presentation"},
    //     //     Todo{Name: "Host meetup"},
    //     // }

    //     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //     w.WriteHeader(http.StatusOK)

    //     if err := json.NewEncoder(w).Encode(todos); err != nil {
    //         panic(err)
    //     }
    // }

    // func TodoShow(w http.ResponseWriter, r *http.Request) {
    //     vars := mux.Vars(r)
    //     todoId := vars["todoId"]
    //     fmt.Fprintln(w, "Todo show:", todoId)
    // }

    // func TodoCreate(w http.ResponseWriter, r *http.Request) {
    //     var todo Todo
    //     body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    //     if err != nil {
    //         panic(err)
    //     }
    //     if err := r.Body.Close(); err != nil {
    //         panic(err)
    //     }
    //     if err := json.Unmarshal(body, &todo); err != nil {
    //         w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //         w.WriteHeader(422) // unprocessable entity
    //         if err := json.NewEncoder(w).Encode(err); err != nil {
    //             panic(err)
    //         }
    //     }

    //     t := RepoCreateTodo(todo)
    //     w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    //     w.WriteHeader(http.StatusCreated)
    //     if err := json.NewEncoder(w).Encode(t); err != nil {
    //         panic(err)
    //     }
    // }
*/


type StudentHandler struct {
    session *mgo.Session
}

func NewStudentHandler(s *mgo.Session) *StudentHandler {
    return &StudentHandler{s}
}

func (h StudentHandler) GetStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    netId := vars["id"]


}

func (h StudentHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
    //create a pointer to the Student struct
    //Student struct maps to the JSON body
    reqBodyStruct := new(models.Student)
    //create a new decoder around the http request body, then feed it the newly created struct variable
    if err := json.NewDecoder(r.Body).Decode(&reqBodyStruct); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // do stuff


    return
}

func (h StudentHandler) RemoveStudent(w http.ResponseWriter, r *http.Request) {

}

func (h StudentHandler) UpdateStudent(w http.ResponseWriter, r *http.Request) {

}

func (h StudentHandler) ListStudents(w http.ResponseWriter, r *http.Request) {

}

func Beer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Drink Yer Beer!")
}
