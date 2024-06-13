package main

import (
	"WebDetection/database"
	handler "WebDetection/handler/home"
	rabbitmq "WebDetection/rabbitMQ"
	repoimpl "WebDetection/repository/repo_impl"
	"WebDetection/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

func main() {
	sql := &database.Sql{
		Host:     "localhost",
		User:     "postgres",
		Port:     5432,
		Password: "postgres",
		Dbname:   "web-detection",
	}
	e := echo.New()
	e.Use(middleware.CORS())
	sql.Connect()
	defer sql.Close()

	db := handler.SiteHandler{
		Repo: repoimpl.NewDataRepo(sql),
	}

	api := router.API{
		Echo:    e,
		Handler: db,
	}
	api.SetUpRouter()
	go func() {
		if err := e.Start(":1910"); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	// Start consuming messages from RabbitMQ
	go func() {
		consumeRepo := rabbitmq.ConsumeRepo{}
		consumeRepo.Receiver(echo.New().AcquireContext(), sql)
	}()

	// Block main goroutine to keep the program running
	select {}
}
