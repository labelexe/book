package main

import (
	"fmt"

	config "github.com/reucot/parser/config/collector"
	service "github.com/reucot/parser/internal/service/collector"
	"github.com/reucot/parser/internal/storage/repository"
	"github.com/reucot/parser/internal/storage/repository/psql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := config.Load(); err != nil {
		logrus.Fatal(fmt.Errorf("failed to load config: %w", err))
	}

	db, err := psql.New(psql.Config{
		Username: config.Get().DB.User,
		Password: config.Get().DB.Password,
		DBName:   config.Get().DB.Name,
		Host:     config.Get().DB.Host,
		Port:     config.Get().DB.Port,
		SSLMode:  config.Get().DB.SslMode,
	})

	if err != nil {
		logrus.Fatal(fmt.Errorf("unable connect to postgres: %s", err))
	}

	repo := repository.NewPsql(db)
	service := service.New(repo)
	service.Run()
}
