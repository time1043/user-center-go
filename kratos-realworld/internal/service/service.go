package service

import (
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/biz"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRealWorldService)

// GreeterService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.UserUsecase
}

// NewGreeterService new a greeter service.
func NewRealWorldService(uc *biz.UserUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}
