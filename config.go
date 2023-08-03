package main

import (
	"time"
)

type config struct {
	UsernameKey  string        `env:"USERNAME_KEY" envDefault:"username"`
	GameDuration time.Duration `env:"GAME_DURATION" envDefault:"30s"`
	HTTPAddr     string        `env:"HTTP_ADDR" envDefault:"0.0.0.0:80"`
	SocTimeout   time.Duration `env:"LISTENCON" envDefault:"15s"`
}
