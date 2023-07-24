package bootstrap

import (
	"WriteWise/pkg/config"
	"WriteWise/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	routing.RegisterRoutes()

	routing.Serve()
}
