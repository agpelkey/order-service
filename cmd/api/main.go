package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/agpelkey/order-service/domain"
	"github.com/agpelkey/order-service/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type config struct {
    port int
    env string
}

type application struct {
    config      config
    UserStore   domain.CustomerService 
    EntreeStore domain.EntreeService
    CartStore domain.CartService
}

func main() {
    
    // create config
    var cfg config

    flag.IntVar(&cfg.port, "port", 8080, "API Server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()

    dbpool, err := pgxpool.New(context.Background(), os.Getenv("FOOD_SERVICE_DSN"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to create connection pool")
        os.Exit(1)
    }

    defer dbpool.Close()

    app := &application{
        config: cfg,
        UserStore: postgres.NewUserStore(dbpool),
        EntreeStore: postgres.NewEntreeStore(dbpool),
        CartStore: postgres.NewCartStore(dbpool),
        
    }
    // start server
    err = app.server()
    if err != nil {
        log.Fatal(err)
    }
    
}
