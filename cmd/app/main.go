package main

import (
	"github.com/itmosha/vk-internship-2024/internal/app"
	"github.com/itmosha/vk-internship-2024/internal/config"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}
