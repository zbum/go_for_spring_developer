//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	bean2 "go_for_spring_developer/10-dependency-injection/e3/bean"
)

func initializeUserService() *bean2.UserService {
	wire.Build(bean2.NewUserService, bean2.NewUserRepository)
	return &bean2.UserService{}
}
