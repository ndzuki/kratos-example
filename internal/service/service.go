package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"

	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealworldService)

// RealworldService is a greeter service.
type RealworldService struct {
	v1.UnimplementedRealworldServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewRealworldService new a greeter service.
func NewRealworldService(uc *biz.UserUsecase, logger log.Logger) *RealworldService {
	return &RealworldService{uc: uc, log: log.NewHelper(logger)}
}
