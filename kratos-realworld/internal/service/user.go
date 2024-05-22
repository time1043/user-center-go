package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.UserReply, error) {
	u, err := s.uc.Login(ctx, in.User.Email, in.User.Password)
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.Username,
			Email:    u.Email,
			Bio:      u.Bio,
			Image:    u.Image,
			// Token:    u.Token,
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}

	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.Username,
			Email:    u.Email,
			Bio:      u.Bio,
			Image:    u.Image,
			Token:    u.Token,
		},
	}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
