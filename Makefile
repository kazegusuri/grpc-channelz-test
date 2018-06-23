
all: pb/chan.pb.go server/server sub/sub call/call

server/server: server/main.go
	go build -o server/server server/main.go

sub/sub: sub/main.go
	go build -o sub/sub sub/main.go

call/call: call/main.go
	go build -o call/call call/main.go

pb/chan.pb.go: pb/chan.proto
	protoc --plugin=protoc-gen-grpc=`which protoc-gen-go` --go_out=plugins=grpc:. pb/chan.proto
