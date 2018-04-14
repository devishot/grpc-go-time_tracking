godep-install:
	go get -u github.com/golang/protobuf/protoc-gen-go

proto-update:
	git submodule update --init --recursive

proto-compile:
	mkdir -p api/
	protoc -I=grpc-protofiles/time_tracking/ -I=${GOPATH}/src --go_out=api/ grpc-protofiles/time_tracking/api.proto

run:
	go run main.go