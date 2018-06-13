# TrackMyFlippinTime
a web application (with the company of a cli script),to track your time accross **devices**(still no support for native mobile)
you can run it as a service anywhere heroku,your PC,A pi, the neighbours unpatched router the possiblites are endless!

## Build Instructions 

install golang

###### linux
```
sudo apt install golang-go #debain
sudo dnf install golang #fedora
```

###### windows
Download and install [the golang MSI](https://golang.org/doc/install?download=go1.10.3.windows-amd64.msi)

### CLI
replace the site variable with whatever host that the service is run on ,make sure it contains http(s) and the port number if necessary 
then

```
go build cli.go #build the CLI script
sudo mv cli /usr/local/bin/trackmft #add it to the $PATH
```

### Service
you can just download a precompiled executable for your platform from the [releases page](https://github.com/Mohamedemad4/TrackMyFlippinTime/releases) 

or you can build it from source:

make a new $GOPATH to build the package
```
mkdir go
export GOPATH=$(PWD)/go
```
install the go-sqlite3 and mux 
```
go get github.com/gorilla/mux
go get github.com/mattn/go-sqlite3
```

finally Build the service

```
go build app.go
```

and execute it 
```
./app
```

## Service API

the service API has 4 endpoints 
####newstatement
```
GET /newstatement/<statement>/<statement_encoded>
```

where Statement is the full statement denoted by statment_encoded in other requests

##### transaltestatement
```
GET /transaltestatement/<statement_encoded>/
```
return the actual statement from an encoded statement (Remeber Statements is how you discibe shunks of your time)
##### deposit
```
GET /deposit/<fromtimestamp>/<totimestamp>/<statement_encode>
```

where fromtimestamp and totimestamp are Unix Timestamps of what he where you doing between the 1st and the 2nd respectivly
and where statement_encode is the encoded statement do denot the activity

##### withdraw
```
GET /withdraw/<fromtimestamp>/<totimestamp>/
```

where fromtimestamp and totimestamp are Unix Timestamps 

and it returns JSON history deposits in the time between fromtimestamp and totimestamp
