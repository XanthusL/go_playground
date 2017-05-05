TIME:=`date +%Y_%m_%d_%H:%M:%S`
app:
	go build -ldflags "-X main.version=${TIME}" app.go
