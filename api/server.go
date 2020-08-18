package api

import (
	"context"
	"fmt"
	"kirby/api/healthcheck"
	"kirby/api/user"
	"kirby/config"
	"kirby/database"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

// Server struct
type Server struct {
	Router *chi.Mux
}

// NewServer returns a fully initialized application struct
func NewServer() *Server {
	config.LoadEnv()
	router := chi.NewRouter()
	application := &Server{router}

	pg, err := database.PgConnect()
	if err != nil {
		log.Fatalf("Postgres connection failed: %v\n", err)
	}
	pg.AutoMigrate(&user.User{})

	redis, err := database.RedisConnect()
	if err != nil {
		log.Fatalf("Redis connection failed: %v\n", err)
	}
	userService := &user.Service{DB: pg, Redis: redis}

	allowedOrigins := []string{"*"}
	cors := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler,
		JwtAuthentication,
	)

	router.Route("/health", func(r chi.Router) {
		r.Get("/server", healthcheck.CheckHTTPConnection)
		r.Get("/database", healthcheck.CheckDatabaseConnection(database.Pg))
	})

	router.Route("/users", func(r chi.Router) {
		r.Get("/{id}", user.FetchUser(userService))
		r.Post("/", user.CreateUser(userService))
	})

	router.Route("/tokens", func(r chi.Router) {
		r.Post("/", user.Authenticate(userService))
		r.Put("/", user.RefreshToken(userService))
	})

	return application
}

// Start running the application
func (app *Server) Start() {

	var runChannel = make(chan os.Signal, 1)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		config.Env.ServerTimeout,
	)
	defer cancel()

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Env.Host, config.Env.Port),
		Handler:      app.Router,
		ReadTimeout:  config.Env.ReadTimeout,
		WriteTimeout: config.Env.WriteTimeout,
		IdleTimeout:  config.Env.IdleTimeout,
	}

	signal.Notify(runChannel, os.Interrupt, syscall.SIGTSTP)

	log.Printf("Server is starting on %s (%s environment)\n", server.Addr, config.Env.Name)

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
