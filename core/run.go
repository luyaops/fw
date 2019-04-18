package core

import (
	"github.com/luyaops/fw/common/config"
	"github.com/luyaops/fw/common/log"
	"google.golang.org/grpc"
	"net"
)

type RpcServer struct {
	Server *grpc.Server
	listen net.Listener
}

// TODO: Interceptors https://github.com/grpc-ecosystem/go-grpc-middleware
func NewRpcServer() *RpcServer {
	server := grpc.NewServer(grpc.UnaryInterceptor(ChainUnaryServer(
	//interceptor.Recovery(),
	//interceptor.DependInject(),
	//interceptor.Authenticate(),
	)))
	return &RpcServer{Server: server}
}

func (s *RpcServer) Run() error {
	listen, err := net.Listen("tcp", config.RpcBind)
	if err != nil {
		return err
	}
	s.listen = listen
	log.Infof("Listening on %v", config.RpcBind)
	return s.Server.Serve(s.listen)
}
