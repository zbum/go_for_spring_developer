//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go_for_spring_developer/01-class/e3/bean"
)

func initializeUserService() *bean.UserService {
	wire.Build(bean.NewUserService, bean.NewUserRepository)
	return &bean.UserService{}
}
