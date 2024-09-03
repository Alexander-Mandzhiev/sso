package auth

import (
	"context"
	"errors"
	"fmt"
	contract "sso/contract/gen/go/sso"
	"sso/internal/domain/entity"
	"sso/internal/repository"
	service "sso/internal/service/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	contract.UnsafeAuthServer
	service *service.Auth
}

func Register(gRPC *grpc.Server, service *service.Auth) {
	contract.RegisterAuthServer(gRPC, &serverAPI{service: service})
}

const (
	emptyValue = 0
)

func (s *serverAPI) Signin(ctx context.Context, req *contract.SigninRequest) (*contract.SigninResponse, error) {

	if err := validateSingin(req); err != nil {
		return nil, err
	}

	signin := &entity.Signin{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		AppID:    int(req.GetAppId()),
	}

	token, err := s.service.Signin(ctx, signin)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect email or password")
		}
		return nil, status.Error(codes.Internal, "Internal error")
	}

	return &contract.SigninResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Signup(ctx context.Context, req *contract.SignupRequest) (*contract.SignupResponse, error) {

	if err := validateSingup(req); err != nil {
		return nil, err
	}

	signup := &entity.Signup{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	userID, err := s.service.Signup(ctx, signup)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, service.ErrUserExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &contract.SignupResponse{
		UserId: userID,
	}, nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *contract.IsAdminRequest) (*contract.IsAdminResponse, error) {

	if err := validateIsAdmin(req); err != nil {
		return nil, err
	}
	isAdmin, err := s.service.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &contract.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
