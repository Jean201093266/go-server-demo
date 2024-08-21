package main

import (
	"context"
	"grpc_demo/client"
	"grpc_demo/kitex_gen/example"
	"grpc_demo/kitex_gen/user_service"
	"grpc_demo/model"
	"grpc_demo/redis"
	"log"
)

// ExampleServiceImpl implements the last service interface defined in the IDL.
type ExampleServiceImpl struct{}

// Ping implements the ExampleServiceImpl interface.
func (s *ExampleServiceImpl) Ping(ctx context.Context, req *example.CommonMsg) (resp *example.CommonMsg, err error) {
	log.Print(req)
	qr, _ := model.QueryDemo()
	log.Println(qr)
	rr := redis.Get("gotest")
	log.Print(rr)
	r, err1 := client.GetUserServiceClient().Ping(context.Background(), &user_service.PingRequest{Service: "jinjin"})
	if err1 != nil {
		log.Print(err1)
	}
	log.Print(r)
	return &example.CommonMsg{Data: r.GetData()}, nil
}
