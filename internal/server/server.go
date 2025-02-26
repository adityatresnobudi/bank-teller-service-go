package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adityatresnobudi/bank-teller-service-go/config"
	serviceHandler "github.com/adityatresnobudi/bank-teller-service-go/internal/domain/service/handler"
	serviceService "github.com/adityatresnobudi/bank-teller-service-go/internal/domain/service/service"
	userHandler "github.com/adityatresnobudi/bank-teller-service-go/internal/domain/user/handler"
	userService "github.com/adityatresnobudi/bank-teller-service-go/internal/domain/user/service"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/service_repo/service_pg"
	"github.com/adityatresnobudi/bank-teller-service-go/internal/repositories/user_repo/user_pg"
	"github.com/adityatresnobudi/bank-teller-service-go/pkg/postgres"
	"github.com/gin-gonic/gin"
)

type server struct {
	cfg config.Config
	r   *gin.Engine
}

func NewServer(cfg config.Config) *server {
	return &server{
		cfg: cfg,
		r:   gin.Default(),
	}
}

func (s *server) Run() {
	db, err := postgres.NewDB(
		s.cfg.Postgres.Host,
		s.cfg.Postgres.Port,
		s.cfg.Postgres.User,
		s.cfg.Postgres.Password,
		s.cfg.Postgres.DBName,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = postgres.InitializeTable(db); err != nil {
		if err = db.Close(); err != nil {
			log.Printf("db graceful shutdown: %s\n", err.Error())
		} else {
			fmt.Printf("db graceful shutdown succeeded\n")
		}
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	userRepo := user_pg.NewUserRepo(db)
	serviceRepo := service_pg.NewServiceRepo(db)

	userService := userService.NewUserService(userRepo)
	serviceService := serviceService.NewServiceService(serviceRepo)

	userHandler := userHandler.NewUserHandler(s.r, ctx, userService)
	serviceHandler := serviceHandler.NewServiceHandler(s.r, ctx, serviceService)

	userHandler.MapRoutes()
	serviceHandler.MapRoutes()

	go func() {
		log.Printf("Listening on PORT: %s\n", s.cfg.Http.Port)
		if err := s.runGinServer(); err != nil {
			log.Printf("s.r.Run: %s\n", err.Error())
		}

	}()

	oscall := <-ch

	if err = db.Close(); err != nil {
		log.Printf("db graceful shutdown: %s\n", err.Error())
	} else {
		fmt.Printf("db graceful shutdown succeeded\n")
	}

	fmt.Printf("system call: %+v\n", oscall)
}
