// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userservices

import (
	server "github.com/cloudwego/kitex/server"
	user_service "grpc_demo/kitex_gen/user_service"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler user_service.UserServices, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
