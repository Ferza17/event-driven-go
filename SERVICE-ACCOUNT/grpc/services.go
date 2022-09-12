package grpc

import (
	"github.com/Ferza17/event-driven-account-service/model/pb"
	userService "github.com/Ferza17/event-driven-account-service/module/user/presenter/grpc"
)

func (srv *Server) RegisterService() {
	// CreateUser Service, Service can be multiple
	pb.RegisterUserServiceServer(srv.grpcServer, userService.NewUserGRPCPresenter())
}
