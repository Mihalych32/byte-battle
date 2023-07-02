package app

import (
	"byte-battle_backend/config"
	"byte-battle_backend/internal/handler"
	"byte-battle_backend/internal/repo"
	"byte-battle_backend/internal/usecase"
	"byte-battle_backend/pkg/postgres"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	pg, err := postgres.NewPostgres(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Pass, cfg.DB.Name, cfg.DB.SSLMode, cfg.DB.MaxIdleConns, cfg.DB.MaxOpenConns)
	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepo := repo.NewUserRepo(pg)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	router := gin.New()

	runport := os.Getenv("PUBLIC_BACKEND_PORT")
	if runport == "" {
		err = fmt.Errorf("PUBLIC_BACKEND_PORT variable not found in .env file")
		return
	}
	runport = ":" + runport
	handler.StartNewServer(router, userHandler, runport)
}
