build-proto:
	protoc -I=./proto --go_out=plugins=grpc:./app/proto proto/*.proto

build-server:
	GO111MODULE=on go build -o exe ./cmd/

evans:
	evans ./proto/server.proto --port 3000 --package proto --service Server