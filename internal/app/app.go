package app

import (
	"log"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/config"
	"github.com/itmosha/vk-internship-2024/internal/http_server"
)

// Create all necessary dependencies and run application.
func Run(cfg *config.Config) {
	// pg, err := postgres.NewPostgres(cfg.DB.Address, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
	// if err != nil {
	// 	log.Fatalf("could not create postgres connection: %s\n", err)
	// }
	// _ = pg

	router := http_server.NewRouter()
	log.Println("starting http server on port", cfg.HTTPServer.RunPort)
	err := http.ListenAndServe(":"+cfg.HTTPServer.RunPort, router)
	if err != nil {
		log.Fatalf("could not start http server: %s\n", err)
	}
}
