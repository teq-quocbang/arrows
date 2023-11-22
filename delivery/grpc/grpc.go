package grpc

import (
	"github.com/teq-quocbang/arrows/proto"
	"github.com/teq-quocbang/arrows/usecase"
)

type TeqService struct {
	proto.UnimplementedTeqServiceServer
	UseCase *usecase.UseCase
}
