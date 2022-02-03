package app

import (
	"fmt"
	"time"

	router "github.com/fasthttp/router"
	"github.com/patrickmn/go-cache"
	user_http "github.com/perlinleo/golang-todo/internal/app/user/delivery"
	user_psql "github.com/perlinleo/golang-todo/internal/app/user/repository"
	user_usecase "github.com/perlinleo/golang-todo/internal/app/user/usecase"

	"github.com/valyala/fasthttp"
)

func Start() error {
	config := NewConfig()

	_, err := NewServer(config)
	if err != nil {
		return err
	}
	ConnPool, err := NewDataBase(config.App.DatabaseURL)
	if err != nil {
		return err
	}

	userCache := cache.New(5*time.Minute, 10*time.Minute)
	userRepository := user_psql.NewUserPSQLRepository(ConnPool, userCache)
	

	userUsecase := user_usecase.NewUserUsecase(userRepository)

	router := router.New()
	user_http.NewUserHandler(router, userUsecase)
	
	
	fmt.Printf("STARTING SERVICE ON PORT %s\n", config.App.Port)
	err = fasthttp.ListenAndServe(config.App.Port, router.Handler)
	if err != nil {
		return err
	}

	return nil
}