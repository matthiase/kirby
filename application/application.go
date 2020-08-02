package application

import (
	"context"
	"fmt"
	"kirby/config"
	"kirby/dbclient"
	"kirby/healthcheck"
	"kirby/redisclient"
	"kirby/user"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Application struct
type Application struct {
	Router *chi.Mux
}

// NewApplication returns a fully initialized application struct
func NewApplication() *Application {
	dbClient, err := dbclient.Connect(config.DatabaseURI())
	if err != nil {
		log.Fatalln("Database connection failed")
	}

	redisClient, err := redisclient.Connect("localhost:6379", "", 0)
	if err != nil {
		log.Fatalln("Redis connection failed")
	}

	dbClient.AutoMigrate(&user.User{})

	userService := &user.Service{DB: dbClient, Redis: redisClient}

	router := chi.NewRouter()
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		JwtAuthentication,
	)

	router.Route("/health", func(r chi.Router) {
		r.Get("/server", healthcheck.CheckHTTPConnection)
		r.Get("/database", healthcheck.CheckDatabaseConnection(dbClient))
	})

	router.Route("/users", func(r chi.Router) {
		r.Get("/{id}", user.FetchUser(userService))
		r.Post("/", user.CreateUser(userService))
	})

	router.Route("/tokens", func(r chi.Router) {
		r.Post("/", user.Authenticate(userService))
		r.Put("/", user.RefreshToken(userService))
	})

	application := &Application{router}
	return application
}

// Start running the application
func (app *Application) Start() {

	var runChannel = make(chan os.Signal, 1)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.ServerTimeout(),
	)
	defer cancel()

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host(), config.Port()),
		Handler:      app.Router,
		ReadTimeout:  config.ReadTimeout(),
		WriteTimeout: config.WriteTimeout(),
		IdleTimeout:  config.IdleTimeout(),
	}

	signal.Notify(runChannel, os.Interrupt, syscall.SIGTSTP)

	log.Printf("Server is starting on %s\n", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatalf("Server failed to start: %v", err)
			}
		}
	}()

	interrupt := <-runChannel

	log.Printf("Server is shutting down gracefully: %v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	}
}
