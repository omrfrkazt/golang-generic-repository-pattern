package main

import (
	"fmt"
	"os"

	"github.com/omrfrkazt/golang-generic-repository-pattern/config"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/app"
)

const dev = "development"
const production = "production"

func main() {
	cfg, err := config.NewConfig(dev)
	if err != nil {
		fmt.Println("Error reading config file", err)
		os.Exit(1)
	}
	app.Run(cfg)
}
