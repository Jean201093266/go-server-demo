// Code generated by Kitex v0.8.0. DO NOT EDIT.

package exampleservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	example "grpc_demo/kitex_gen/example"
)

func serviceInfo() *kitex.ServiceInfo {
	return exampleServiceServiceInfo
}

var exampleServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ExampleService"
	handlerType := (*example.ExampleService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Ping": kitex.NewMethodInfo(pingHandler, newPingArgs, newPingResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "example",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func pingHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(example.CommonMsg)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(example.ExampleService).Ping(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PingArgs:
		success, err := handler.(example.ExampleService).Ping(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PingResult)
		realResult.Success = success
	}
	return nil
}
func newPingArgs() interface{} {
	return &PingArgs{}
}

func newPingResult() interface{} {
	return &PingResult{}
}

type PingArgs struct {
	Req *example.CommonMsg
}

func (p *PingArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(example.CommonMsg)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PingArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PingArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PingArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PingArgs) Unmarshal(in []byte) error {
	msg := new(example.CommonMsg)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PingArgs_Req_DEFAULT *example.CommonMsg

func (p *PingArgs) GetReq() *example.CommonMsg {
	if !p.IsSetReq() {
		return PingArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PingArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PingArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PingResult struct {
	Success *example.CommonMsg
}

var PingResult_Success_DEFAULT *example.CommonMsg

func (p *PingResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(example.CommonMsg)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PingResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PingResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PingResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PingResult) Unmarshal(in []byte) error {
	msg := new(example.CommonMsg)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PingResult) GetSuccess() *example.CommonMsg {
	if !p.IsSetSuccess() {
		return PingResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PingResult) SetSuccess(x interface{}) {
	p.Success = x.(*example.CommonMsg)
}

func (p *PingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PingResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Ping(ctx context.Context, Req *example.CommonMsg) (r *example.CommonMsg, err error) {
	var _args PingArgs
	_args.Req = Req
	var _result PingResult
	if err = p.c.Call(ctx, "Ping", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
