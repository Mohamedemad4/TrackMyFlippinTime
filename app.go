package main

import (
    "net/http"
    "log"
    "fmt"
    "os"
    "database/sql"
    "io/ioutil"
    "encoding/json"
    _ "github.com/mattn/go-sqlite3"
    "github.com/gorilla/mux"
)

var db, err = sql.Open("sqlite3", "./tmft_service.db")

func init(){
    if err!=nil{
        log.Fatal(err)
    }
    if _, err := os.Stat("Ang.js"); os.IsNotExist(err){
        response, err:= http.Get("https://raw.githubusercontent.com/Mohamedemad4/TrackMyFlippinTime/master/Ang.js")
        if err!=nil{log.Fatal(err)}
        body, err := ioutil.ReadAll(response.Body)
        if err!=nil{log.Fatal(err)}
        ioutil.WriteFile("Ang.js", body, 0644)
    }
    if _, err := os.Stat("index.html"); os.IsNotExist(err){
        response, err:= http.Get("https://raw.githubusercontent.com/Mohamedemad4/TrackMyFlippinTime/master/index.html")
        if err!=nil{log.Fatal(err)}
        body, err := ioutil.ReadAll(response.Body)
        if err!=nil{log.Fatal(err)}
        ioutil.WriteFile("index.html", body, 0644)
    }  
    _, err := db.Query("SELECT * FROM statements LIMIT 1;")
    if err != nil {
        sqlStmt := `
        CREATE TABLE statements (statement VARCHAR(50),statement_encoded VARCHAR(50));
        CREATE TABLE history (fromstamp int(8), tostamp int(8),statement_encoded VARCHAR(59));
        ` 
        _, err = db.Exec(sqlStmt)
        if err != nil {
            log.Printf("%q: %s\n", err, sqlStmt)
        }
    }
}
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", rootPage).Methods("GET")
    router.HandleFunc("/Ang.js", JsPage).Methods("GET")
    router.HandleFunc("/withdraw/{startstamp}/{endstamp}",withdraw).Methods("GET")
    router.HandleFunc("/deposit/{from}/{to}/{statement_encoded}",deposit).Methods("GET")
    router.HandleFunc("/newstatement/{statement}/{statement_encoded}",newstatement).Methods("GET")
    router.HandleFunc("/transaltestatement/{statement_encoded}",transalteStatement).Methods("GET")
    log.Println("Started Server on :9922")
   
    defer db.Close()
    log.Fatal(http.ListenAndServe(":9922", router))
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	file,err:=ioutil.ReadFile("index.html")
	if err!=nil{
	  log.Println("Error reading index.html",err)
	}
	fmt.Fprintf(w,string(file))
}

func JsPage(w http.ResponseWriter, r *http.Request) {
    file,err:=ioutil.ReadFile("Ang.js")
    if err!=nil{
      log.Println("Error reading Ang.js",err)
    }
    fmt.Fprintf(w,string(file))
}

func withdraw(w http.ResponseWriter,r *http.Request){
    params := mux.Vars(r)
    start := params["startstamp"]
    end := params["endstamp"]

    var historySlice []interface{}

    rows, err := db.Query("SELECT fromstamp,tostamp,statement_encoded from history WHERE fromstamp BETWEEN ? AND ? OR tostamp BETWEEN ? AND ?",
        start,end,start,end)

    if err != nil {
        log.Println(err)
    }else{
       defer rows.Close()
       for rows.Next() {
           var fromstamp int
           var tostamp int
           var statement_encoded string
           err = rows.Scan(&fromstamp,&tostamp,&statement_encoded)
           if err != nil {
               log.Println(err)
           }
           historySlice=append(historySlice,map[string]interface{}{"fromstamp":
            fromstamp, "tostamp": tostamp,"statement_encoded":statement_encoded})
       }
       err = rows.Err()
       if err != nil {
           log.Println(err)
       }
    }
    json.NewEncoder(w).Encode(historySlice)  
}

func deposit(w http.ResponseWriter,r *http.Request){
   
    params := mux.Vars(r)
    
    from:=params["from"]
    to:=params["to"]
    statement_encoded:=params["statement_encoded"]
   
    insertStat , err := db.Prepare("INSERT INTO history VALUES (?,?,?)")
    if err != nil {
        log.Println(err)
        fmt.Fprintf(w,"+ERROR")
    }else{
        insertStat.Exec(from,to,statement_encoded)
    }

}

func transalteStatement(w http.ResponseWriter, r *http.Request){
    params:=mux.Vars(r)
    statement_encoded:=params["statement_encoded"]
    rows,err := db.Query("SELECT statement from statements WHERE statement_encoded=?",statement_encoded)
    if err!=nil{
        log.Println(err)
        fmt.Fprintf(w,"ERROR")
    }
    defer rows.Close()
    rows.Next()
    var statement string
    err = rows.Scan(&statement)
    if err != nil {
        log.Println(err)
    }
    translation:=map[string]string{"statement":statement}
    json.NewEncoder(w).Encode(translation)
    
}

func newstatement(w http.ResponseWriter,r *http.Request){
    params:= mux.Vars(r)
    statement:=params["statement"]
    statement_encoded:=params["statement_encoded"]
    insertStat, err := db.Prepare("INSERT INTO statements VALUES (?,?)")
    if err != nil {
        log.Println(err)
        fmt.Fprintf(w,"ERROR")
    }else{
        insertStat.Exec(statement, statement_encoded)
    }
}