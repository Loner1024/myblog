//go:build wireinject
// +build wireinject

package main

import (
	"github.com/Loner1024/uniix.io/configs"
	"github.com/Loner1024/uniix.io/internal/domain/blog"
	"github.com/Loner1024/uniix.io/internal/server"
	"github.com/Loner1024/uniix.io/internal/services"
	"github.com/Loner1024/uniix.io/internal/store/firebase"
	"github.com/google/wire"
)

func wireApp(configs.Config) (*server.Server, error) {
	panic(wire.Build(server.ProviderSet, blog.ProviderSet, services.ProviderSet, firebase.NewStore))
}
