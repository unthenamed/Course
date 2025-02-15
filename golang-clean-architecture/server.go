package main

import (
	"api-laundy/config"
	"api-laundy/controller"
	"api-laundy/repo"
	"api-laundy/service"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	transactionService service.TransactionService
	engine             *gin.Engine
	host               string
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/")
	controller.ConstructorTransactionController(s.transactionService, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}

}
func initDB(cfg *config.Config) *sql.DB {
	sqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require",
		cfg.DBConfig.User,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.Database,
	)

	db, err := sql.Open(
		cfg.DBConfig.Driver,
		sqlInfo,
	)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	db := initDB(cfg)
	host := fmt.Sprintf(":%s", cfg.APIConfig.Port)

	transaction := service.NewTransactionService(
		repo.NewTransaction(
			db,
		),
	)

	return &Server{
		host:               host,
		transactionService: transaction,
		engine:             gin.Default(),
	}
}
