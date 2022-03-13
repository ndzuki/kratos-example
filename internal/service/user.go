package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealworldService) Login(ctx context.Context, req *v1.LoginRequest) (rep *v1.UserReply, err error) {
	rv, err := s.uc.Login(ctx, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: rv.Username,
			Token:    rv.Token,
		},
	}, nil
}

func (s *RealworldService) Register(ctx context.Context, req *v1.RegisterRequest) (rep *v1.UserReply, err error) {
	u, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: u.Username,
			Email:    u.Email,
			Bio:      "something",
			Token:    u.Token,
			Image:    "someimage",
		},
	}, nil
}

func (s *RealworldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (rep *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}

func (s *RealworldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (rep *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}

func (s *RealworldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (rep *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
