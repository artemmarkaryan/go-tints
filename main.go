package main

import (
	"github.com/artemmarkaryan/gotints/config"
	"github.com/artemmarkaryan/gotints/internal/controller"
	)

func main() {
	config.ConfigureProject()
	controller.Serve()
}
