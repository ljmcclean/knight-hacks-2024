package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	_ "github.com/lib/pq"

	"github.com/ljmcclean/knight-hacks-2024/config"
	"github.com/ljmcclean/knight-hacks-2024/postgres"
	"github.com/ljmcclean/knight-hacks-2024/seeds"
	"github.com/ljmcclean/knight-hacks-2024/server"
)

func main() {
	ctx := context.Background()
	err := run(ctx)
	if err != nil {
		log.Fatalf("error closing server: %s", err)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg := config.New()
	cfg.Init()

	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatalf("error connecting to database: %s\n", err)
	}
	defer db.Close()
	log.Printf("established database connection\n")

	server := server.New(cfg, ctx, db)

	seeds.SeedProjects(ctx, db)

	go func() {
		log.Printf("listening and serving on %s\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, cfg.Server.KillTime*time.Second)
		defer cancel()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("error shutting down server: %s\n", err)
		}
	}()
	wg.Wait()

	return nil
}
