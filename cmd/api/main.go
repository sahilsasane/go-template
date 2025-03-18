package main

import (
	"flag"
	"os"
	"sync"
)

var version string

type config struct {
	port int
	env  string
	cors struct {
		trustedOrigins []string
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	// models data.Model
	wg sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	// flag.StringVar(&cfg.db.dsn, "db-dsn", "", "SQL DSN")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		config: cfg,
		logger: logger,
		// models: data.NewModels(db),
	}

	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
