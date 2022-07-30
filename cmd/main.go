package main

import (
	"log"

	todo "github.com/Brotiger/todo_go"
	"github.com/Brotiger/todo_go/pkg/handler"
	"github.com/Brotiger/todo_go/pkg/repository"
	"github.com/Brotiger/todo_go/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepositor()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}