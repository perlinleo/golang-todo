package main

import (
	app "github.com/perlinleo/golang-todo/internal/app"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Error().Msgf(app.Start().Error())
}