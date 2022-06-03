package server

import (
	v1 "github.com/Loner1024/uniix.io/api/gen/go/api"
	"github.com/Loner1024/uniix.io/configs"
	"github.com/Loner1024/uniix.io/internal/services"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	conf    configs.Config
	service *services.Service
	logger  *zap.SugaredLogger
}

func NewServer(s *services.Service, conf configs.Config, l *zap.SugaredLogger) *Server {
	return &Server{conf: conf, service: s, logger: l}
}

func (s *Server) Start() error {
	cc, err := net.Listen("tcp", s.conf.Addr+":"+s.conf.Port)
	if err != nil {
		return err
	}
	interceptor := grpc.UnaryInterceptor(
		grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(s.logger.Desugar()),
		),
	)
	srv := grpc.NewServer(interceptor)
	v1.RegisterBlogServiceServer(srv, s.service)
	s.logger.Infof("Start grpc server at: %s:%s", s.conf.Addr, s.conf.Port)
	return srv.Serve(cc)
}
