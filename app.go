package main

import (
    "net/http"
    "log"
    "fmt"
    "github.com/gorilla/mux"
)

func rootPage(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.html",nil)
    t.Execute(w)
}

func withdraw(w http.ResponseWriter,r *http.Request){
    params := mux.Vars(r)
    start := params["startstamp"]
    end := params["endstamp"]
    
}

func deposit(w http.ResponseWriter,r *http.Request){
    params := mux.Vars(r)
    from:=params["from"]
    to:=params["to"]
    statement_encoded:=params["statement_encoded"]
}

func newstatement(w http.ResponseWriter,r *http.Request){
    params:= mux.Vars(r)
    statement:=params["statement"]
    statement_encoded:=params["statement_encoded"]
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", rootPage).Methods("GET")
    router.HandleFunc("/withdraw/{startstamp}/{endstamp}",withdraw).Methods("GET")
    router.HandleFunc("/deposit/{from}/{to}/{statement_encoded}",deposit).Methods("GET")
    router.HandleFunc("/newstatement/{statement}/{statement_encoded}",newstatement).Methods("GET")
    log.Println("Started Server")
    log.Fatal(http.ListenAndServe(":8080", router))
}