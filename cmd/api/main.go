package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"

	"metaverse.sahilsasane.net/internal/data"
	"metaverse.sahilsasane.net/internal/jsonlog"
	"metaverse.sahilsasane.net/internal/mailer"
)

// var version string

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	// cors struct {
	// 	trustedOrigins []string
	// }
	// limiter struct {
	// 	rps     float64
	// 	burst   int
	// 	enabled bool
	// }
	jwt struct {
		secret string
	}
}

type application struct {
	config config
	models data.Models
	logger *jsonlog.Logger
	mailer mailer.Mailer
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://metaverse:password@localhost:5432/metaverse?sslmode=disable", "postgreSQL SQL DSN")

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "e6e586b8766161", "SMTP username")

	flag.StringVar(&cfg.smtp.password, "smtp-password", "02570437a73b2f", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Metaverse<no-reply@metaverse.sahil.net>", "SMTP sender")

	flag.StringVar(&cfg.jwt.secret, "jwt-secret", "", "JWT secret")

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	// Add connection attempt log
	logger.PrintInfo("attempting database connection", map[string]string{
		"dsn": cfg.db.dsn,
	})

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, map[string]string{
			"dsn": cfg.db.dsn,
		})
	}

	// Log successful connection
	logger.PrintInfo("database connection pool established", nil)

	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	// Log the connection attempt
	fmt.Printf("Attempting to connect to database with DSN: %s\n", cfg.db.dsn)

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	duration, err := time.ParseDuration(cfg.db.maxIdleTime)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v (check if PostgreSQL is running and accessible)", err)
	}

	return db, nil
}
