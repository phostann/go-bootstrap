package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v4"

	"shopping-mono/app/controllers"
	"shopping-mono/app/services"
	"shopping-mono/pkg/configs"
	"shopping-mono/pkg/middlewares"
	"shopping-mono/pkg/routes"
	"shopping-mono/platform/database/postgres"
)

type CleanTask = func()

type App struct {
	Config     *configs.Config
	app        *fiber.App
	cleanTasks []CleanTask
}

// NewApp creates a new App
func NewApp() *App {
	app := fiber.New()
	return &App{
		app: app,
	}
}

func (a *App) start() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = a.app.Shutdown()
	}()
	if err := a.app.Listen(fmt.Sprintf("%s:%s", a.Config.Server.Host, a.Config.Server.Port)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	fmt.Println("Running cleanup tasks")
	wg := &sync.WaitGroup{}
	for _, t := range a.cleanTasks {
		wg.Add(1)
		go func(f CleanTask) {
			f()
			wg.Done()
		}(t)
	}
	wg.Wait()
}

func (a *App) addCleanTask(f CleanTask) {
	a.cleanTasks = append(a.cleanTasks, f)
}

func (a *App) Prepare() *App {
	cfg, err := configs.ParseConfig()
	if err != nil {
		panic(err)
	}
	a.Config = cfg

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", a.Config.Database.User, a.Config.Database.Password, a.Config.Database.Host, a.Config.Database.Port, a.Config.Database.DB)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	a.addCleanTask(func() {
		_ = conn.Close(context.Background())
	})

	queries := postgres.New(conn)
	service := services.NewService(queries)
	controller := controllers.NewController(service, cfg)
	middleware := middlewares.NewMiddleware(cfg)

	routes.SetupRoutes(controller, a.app, middleware)

	a.app.Use(recover.New())
	a.app.Use(logger.New())

	return a
}

func (a *App) Run() {
	v1 := a.app.Group("v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World!",
		})
	})
	a.start()
}
