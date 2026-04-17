package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/thatguy-git/ecom/internal/env"
)

func main() {
	ctx := context.Background()
	cfg := config{
		addr: ":1400",
		db: dbConfig{
			dsn: env.GetString("DB_STRING", "host=localhost user=postgres password=postgress dbname=postgres sslmode=disable"),
		}}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)

	}
	defer conn.Close(ctx)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	logger.Info("connected to database", "db string", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		log.Printf("error starting server: %v", err)
		os.Exit(1)
	}
}
