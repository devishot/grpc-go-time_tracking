godep-install:
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get google.golang.org/grpc
	go get github.com/lib/pq
	go get github.com/BurntSushi/toml
	go get github.com/satori/go.uuid

proto-update:
	git submodule update --init --recursive

proto-compile:
	mkdir -p interface/api/
	protoc -I=interface/grpc-protofiles/time_tracking/ -I=${GOPATH}/src --go_out=plugins=grpc:interface/api/ api.proto

run:
	go run main.go

test:
	npm test
