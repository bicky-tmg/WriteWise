package bootstrap

import (
	"WriteWise/pkg/config"
	"WriteWise/pkg/html"
	"WriteWise/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
