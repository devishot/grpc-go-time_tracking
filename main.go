package main

import (
	"log"

	"github.com/golang/protobuf/proto"

	"github.com/devishot/grpc-go-time_tracking/api"
)

func main() {
	r := &api.TimeRecord{}
	r.Amount = 30 // 30 minutes = 0.5 hour
	r.Description = "Read laws"

	out, err := proto.Marshal(r)
	if err != nil {
		log.Fatalln("Failed to encode time record:", err)
		return
	}

	r2 := &api.TimeRecord{}
	err2 := proto.Unmarshal(out, r2)
	if err2 != nil {
		log.Fatalln("Failed to decode time record:", err2)
	}

	log.Println("decoded data:", r2.String())
}
