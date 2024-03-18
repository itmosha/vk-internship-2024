package main

import (
	"github.com/itmosha/vk-internship-2024/internal/app"
	"github.com/itmosha/vk-internship-2024/internal/config"
)

// @Version 1.0.0
// @Title Film Library API
// @Description A simple REST API for working with films and actors data.
// @Security Authorization
// @SecurityScheme JWT http bearer Your JWT token
func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}
