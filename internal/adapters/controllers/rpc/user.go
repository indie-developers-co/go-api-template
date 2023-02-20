package rpc

import (
	"context"

	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/models/request"
	"gitlab.com/indie-developers/go-api-echo-template/internal/domains/repositories"
	"gitlab.com/indie-developers/go-api-echo-template/pb"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userController struct {
	pb.UnimplementedUserServiceServer
	userUseCases repositories.UserUseCases
	validator    validator.Validator
}

func NewUserController(userUseCases repositories.UserUseCases, validator validator.Validator) pb.UserServiceServer {
	return &userController{validator: validator, userUseCases: userUseCases}
}

func (u *userController) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	model := request.CreateUserRequest{
		Name:     in.GetName(),
		LastName: in.GetLastName(),
		Email:    in.GetEmail(),
	}

	if err := u.validator.Validate(&model); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	if err := u.userUseCases.Create(ctx, model); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreateUserResponse{}, nil
}
