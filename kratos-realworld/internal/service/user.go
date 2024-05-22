package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "admin",
			Email:    "admin@admin.com",
			Bio:      "I am a admin",
			Image:    "https://example.com/avatar.png",
			Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODssw5c",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: "admin",
		},
	}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
