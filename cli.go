package main

import (
"fmt"
"strconv"
"os"
"strings"
"log"
"time"
"net/http"
"io/ioutil"
)

func main(){
    site:="http://localhost:9922"
    fromstring:=os.Args[1]
    tostring:=os.Args[2]
    statment:=os.Args[3]

    from:=strings.Split(fromstring,":")
    to:=strings.Split(tostring,":")
    
    year,month, day := time.Now().Date()

    fromstamp:=time.Date(year, month, day, toint(from[0]), toint(from[1]), 0, 0, time.UTC).Unix()
    tostamp:=time.Date(year, month, day, toint(to[0]), toint(to[1]), 0, 0, time.UTC).Unix()
    
    resp,err:=http.Get(fmt.Sprintf("%s/deposit/%d/%d/%s",site,fromstamp,tostamp,statment))
    if err !=nil{
        log.Fatal(err)
    }
    
    body, err := ioutil.ReadAll(resp.Body)
    if err !=nil{
        log.Fatal(err)
    }
    api_resp:=string(body)
    if api_resp=="ERROR"{
        log.Fatal("Error Saving time Check the Service Logs")
    }

}

func toint(strint string)int{
    b10,err:=strconv.ParseInt(strint, 10, 32)
    if err!=nil{
        log.Fatal(err)
    }
    return int(b10)
}
