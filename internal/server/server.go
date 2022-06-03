package server

import (
	v1 "github.com/Loner1024/uniix.io/api/gen/go/api"
	"github.com/Loner1024/uniix.io/configs"
	"github.com/Loner1024/uniix.io/internal/services"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	conf    configs.Config
	service *services.Service
}

func NewServer(s *services.Service, conf configs.Config) *Server {
	return &Server{conf: conf, service: s}
}

func (s *Server) Start() error {
	cc, err := net.Listen("tcp", s.conf.Addr+":"+s.conf.Port)
	if err != nil {
		return err
	}
	srv := grpc.NewServer()
	v1.RegisterBlogServiceServer(srv, s.service)
	return srv.Serve(cc)
}
