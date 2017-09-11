package main

import (
    "encoding/json"
    "net/http"
    "time"
    "regexp"
    "io"
    "io/ioutil"

    "github.com/go-pg/pg"
)

var IsAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

var db = pg.Connect(&pg.Options{
    User: "plugdj",
    Password: "plugdj",
    Database: "plugdj",
    Addr: "192.168.50.13:5432",
})

func Now(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(Time{CurrentTime: start.Format("2006-01-02 15:04")}); err != nil {
        panic(err)
    }

}

func NameCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    var later Later
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &later); err != nil {
        w.WriteHeader(422)
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    if !IsAlphanumeric(later.Name) {
        w.WriteHeader(402)
        if err := json.NewEncoder(w).Encode(Error{Error: "invalid name", Payload: later.Name}); err != nil {
            panic(err)
        }
    } else if err := db.Insert(&later); err != nil {
        w.WriteHeader(403)
        json.NewEncoder(w).Encode(Error{Error: "Can not insert date into the DB", Payload: err.Error()})
    } else {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(Status{Status: "Ok"})
    }
}

func Check(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    var laters []Later
    err := db.Model(&laters).Select()
    if err != nil {
        w.WriteHeader(403)
        json.NewEncoder(w).Encode(Error{Error: "Can not get data from DB", Payload: err.Error()})
    }else if laters == nil {
        json.NewEncoder(w).Encode(Status{Status: "There is no data in the db"})
    }else {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(laters)
    }

}

func Invalid(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(401)

    if err := json.NewEncoder(w).Encode(Error{Error: "invalid method"}); err != nil {
        panic(err)
    }

}
