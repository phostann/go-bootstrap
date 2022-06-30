package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"shopping-mono/app/controllers"
	"shopping-mono/app/services"
	"shopping-mono/pkg/configs"
	"shopping-mono/pkg/middlewares"
	"shopping-mono/pkg/routes"
	"shopping-mono/platform/cache/redis"
	"shopping-mono/platform/database/mongodb"
	"shopping-mono/platform/database/mysql"
)

type CleanTask = func()

type App struct {
	Config     configs.Config
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
	cfg, err := configs.Parse()
	if err != nil {
		panic(err)
	}
	a.Config = cfg

	a.app.Use(recover.New())
	a.app.Use(logger.New())
	a.app.Use(cors.New())

	// mysql
	q, closeDBConn := mysql.New(a.Config)
	a.addCleanTask(closeDBConn)

	// mongodb
	_, closeMongodb := mongodb.New(a.Config)
	a.addCleanTask(closeMongodb)

	// redis
	_, closeRedis := redis.New(a.Config)
	a.addCleanTask(closeRedis)

	service := services.NewService(q)
	controller := controllers.NewController(service, a.Config)
	middleware := middlewares.NewMiddleware(a.Config)

	routes.SetupRoutes(controller, a.app, middleware)

	return a
}

func (a *App) Run() {
	a.start()
}
