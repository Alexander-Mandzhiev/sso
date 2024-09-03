package auth

import (
	contract "sso/contract/gen/go/sso"
	"sso/pkg/validate"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateSingin(req *contract.SigninRequest) error {
	if !validate.IsValidEmail(req.GetEmail()) {
		return status.Error(codes.InvalidArgument, "invalid email")
	}
	if !validate.IsValidPassword(req.GetPassword()) {
		return status.Error(codes.InvalidArgument, "the password must be longer than 7 characters and contain lowercase letters, uppercase letters, numbers and special characters")
	}
	if req.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}
	return nil
}

func validateSingup(req *contract.SignupRequest) error {
	if !validate.IsValidEmail(req.GetEmail()) {
		return status.Error(codes.InvalidArgument, "invalid email")
	}
	if !validate.IsValidPassword(req.GetPassword()) {
		return status.Error(codes.InvalidArgument, "the password must be longer than 7 characters and contain lowercase letters, uppercase letters, numbers and special characters")
	}
	return nil
}

func validateIsAdmin(req *contract.IsAdminRequest) error {
	if req.GetUserId() == "" {
		return status.Error(codes.InvalidArgument, "invalid user_id")
	}

	return nil
}
