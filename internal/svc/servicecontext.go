package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"go-zero-token/internal/config"
	"go-zero-token/internal/middleware"
)

type ServiceContext struct {
	Config      config.Config
	TokenMiddle rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		TokenMiddle: middleware.NewTokenMiddleMiddleware().Handle,
	}
}
