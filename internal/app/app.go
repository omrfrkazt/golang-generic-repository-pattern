package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/omrfrkazt/golang-generic-repository-pattern/config"
	"github.com/omrfrkazt/golang-generic-repository-pattern/internal/db"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	config *config.Config
}

func Run(cfg *config.Config) {
	app := new(App)
	app.config = cfg
	app.db = db.Connect(cfg)
	db.MigrateDB(app.db)
	server := fiber.New()
	server.Use(limiter.New(limiter.Config{Max: 100}), cors.New(cors.ConfigDefault), logger.New(CustomLogger()))
	app.Routes(server)
	server.Listen(app.config.Port)
}
