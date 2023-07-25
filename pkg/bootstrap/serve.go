package bootstrap

import (
	"WriteWise/pkg/config"
	"WriteWise/pkg/html"
	"WriteWise/pkg/routing"
	"WriteWise/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
