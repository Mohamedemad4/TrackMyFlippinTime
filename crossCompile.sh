GOOS=linux GOARCH=amd64 go build  -o Service-amd64-linux app.go
GOOS=linux GOARCH=386 go build  -o Service-i386-linux app.go
GOOS=windows GOARCH=386 go build  -o Service-i386-Windows.exe app.go
GOOS=windows GOARCH=amd64 go build  -o Service-amd64-Windows.exe app.go
GOOS=darwin GOARCH=amd64 go build  -o Service-amd64-MacOs app.go
GOOS=linux GOARCH=arm GOARM=7 go build  -o Service-armv7-linux app.go