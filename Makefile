
all: pb/chan.pb.go

server/server: server/main.go
	go build server/main.go -o server/main

pb/chan.pb.go: pb/chan.proto
	protoc --plugin=protoc-gen-grpc=`which protoc-gen-go` --go_out=plugins=grpc:. pb/chan.proto
