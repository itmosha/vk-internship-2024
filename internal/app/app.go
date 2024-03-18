package app

import (
	"log"
	"net/http"

	"github.com/itmosha/vk-internship-2024/internal/config"
	"github.com/itmosha/vk-internship-2024/internal/handler"
	"github.com/itmosha/vk-internship-2024/internal/http_server"
	repo "github.com/itmosha/vk-internship-2024/internal/repo/postgres"
	"github.com/itmosha/vk-internship-2024/internal/usecase"
	"github.com/itmosha/vk-internship-2024/pkg/logger"
	"github.com/itmosha/vk-internship-2024/pkg/postgres"
)

// Create all necessary dependencies and run application.
func Run(cfg *config.Config) {
	// Setup postgres connection
	pg, err := postgres.NewPostgres(cfg.DB.Address, cfg.DB.User, cfg.DB.Password, cfg.DB.Name)
	if err != nil {
		log.Fatalf("could not create postgres connection: %s\n", err)
	}

	// Setup logger
	logger := logger.NewLogger("logs/logs.txt", cfg.Env)

	// Create repos
	filmRepo := repo.NewFilmRepoPostgres(pg)
	actorRepo := repo.NewActorRepoPostgres(pg)
	filmsActorsRepo := repo.NewFilmsActorsRepoPostgres(pg)
	userRepo := repo.NewUserRepoPostgres(pg)

	// Create usecases
	filmUsecase := usecase.NewFilmUsecase(filmRepo, actorRepo, filmsActorsRepo)
	actorUsecase := usecase.NewActorUsecase(actorRepo, filmsActorsRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Create handlers
	filmHandler := handler.NewFilmHander(filmUsecase, logger)
	actorHandler := handler.NewActorHandler(actorUsecase, logger)
	userHandler := handler.NewUserHandler(userUsecase, logger)

	// Setup router
	router := http_server.NewRouter(filmHandler, actorHandler, userHandler)

	// Run server
	s := &http.Server{
		Addr:           ":" + cfg.HTTPServer.RunPort,
		Handler:        router,
		ReadTimeout:    cfg.HTTPServer.Timeout,
		WriteTimeout:   cfg.HTTPServer.Timeout,
		IdleTimeout:    cfg.HTTPServer.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("starting http server on port %s", cfg.HTTPServer.RunPort)
	log.Fatal(s.ListenAndServe())
}
