package main

import (
	"log"

	"github.com/stephanvebrian/expense-manager/account-service/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("error when initializing config")
	}

	err = startApp(cfg)
	if err != nil {
		log.Fatal("error when trying start the application")
	}
}
