package http

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rodkevich/gmp-tickets/lib/db"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/rodkevich/gmp-tickets/internal/configs"
	"github.com/rodkevich/gmp-tickets/lib/validation"
)

type Server struct {
	cfg        *configs.Configs
	database   *pgxpool.Pool
	router     *chi.Mux
	httpServer *http.Server
	validator  *validator.Validate
	logger     *zap.Logger

	// ticketRepo ticket.Repository
	// userRepo   user.Repository
}

func (srv *Server) Initialize() {
	srv.newConfiguration()
	srv.newLogger()
	srv.newDataBase()
	srv.newValidator()
	srv.newRouter()
	srv.initRoutes()
	// srv.setMiddlewares()
}

func NewServer(version string) *Server {
	log.Printf("Ticket service API version: %s\n", version)
	return &Server{}
}
func (srv *Server) Run() error {
	srv.httpServer = &http.Server{
		Addr:           srv.cfg.Api.Host.String() + ":" + srv.cfg.Api.Port,
		Handler:        srv.router,
		ReadTimeout:    srv.cfg.Api.ReadTimeout,
		WriteTimeout:   srv.cfg.Api.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		var prefix string
		if srv.cfg.Api.Name == "dev" {
			prefix = "http://"
		}
		log.Printf("Serving at "+prefix+"%s:%s\n", srv.cfg.Api.Host, srv.cfg.Api.Port)
		showAllRoutes(srv.router)
		err := srv.httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(
		context.Background(),
		srv.cfg.Api.IdleTimeout*time.Second,
	)

	defer shutdown()

	return srv.httpServer.Shutdown(ctx)
}

func (srv *Server) GetConfig() *configs.Configs {
	return srv.cfg
}

func (srv *Server) newConfiguration() {
	srv.cfg = configs.New()
}

func (srv *Server) newRouter() {
	srv.router = chi.NewRouter()
	srv.router.Use(middleware.RequestID)
	srv.router.Use(middleware.Logger)
	srv.router.Use(middleware.Recoverer)
}

func (srv *Server) newLogger() {

	rawJSONConfig := []byte(`{
      "level": "info",
      "encoding": "console",
      "outputPaths": ["stdout", "/tmp/logs"],
      "errorOutputPaths": ["/tmp/errorlogs"],
      "initialFields": {"initFieldKey": "fieldValue"},
      "encoderConfig": {
        "messageKey": "message",
        "levelKey": "level",
        "nameKey": "logger",
        "timeKey": "time",
        "callerKey": "logger",
        "stacktraceKey": "stacktrace",
        "callstackKey": "callstack",
        "errorKey": "error",
        "timeEncoder": "iso8601",
        "fileKey": "file",
        "levelEncoder": "capitalColor",
        "durationEncoder": "second",
        "callerEncoder": "full",
        "nameEncoder": "full",
        "sampling": {
            "initial": "3",
            "thereafter": "10"
        }
      }
    }`)

	config := zap.Config{}
	if err := json.Unmarshal(rawJSONConfig, &config); err != nil {
		log.Fatalf(err.Error())
	}
	var err error
	srv.logger, err = config.Build()
	if err != nil {
		log.Fatalf(err.Error())
	}
	// srv.logger.Debug("This is a DEBUG message")
	srv.logger.Info("This should have an ISO8601 based time stamp",
		zap.Int("hello world", 3),
		zap.Time("time test", time.Now()),
	)
	// srv.logger.Warn("This is a WARN message")
	// srv.logger.Error("This is an ERROR message")
}

func (srv *Server) newDataBase() {

	switch srv.cfg.Database.Driver {
	case "postgres":
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		pool, err := db.NewConnectionPool(ctx, srv.cfg)
		if err != nil {
			log.Fatalf(err.Error())
		}
		srv.database = pool

	case "mysql":
		log.Fatal("Mysql not implemented yet.You must choose a valid database driver")
	default:
		log.Fatal("You must choose a valid database driver")
	}
}

func showAllRoutes(router *chi.Mux) {
	walkFunc := func(method string, path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("path: %s method: %s ", path, method)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Print(err)
	}
}

func (srv *Server) newValidator() {
	srv.validator = validation.NewValidator()
}
