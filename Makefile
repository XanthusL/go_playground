TIME:=`date +%Y_%m_%d_%H:%M:%S`
app:
	go build -ldflags "-X main.version=${TIME}" app.go
server:archive/file_server/file_server.go
	go build archive/file_server/file_server.go
	#(cd archive/file_server/ && go build)