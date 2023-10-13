package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/syumai/workers/cloudflare/d1"
	_ "github.com/syumai/workers/cloudflare/d1"
)

var instance *sql.DB

func getInstance() *sql.DB {
	if instance == nil {
		con, err := d1.OpenConnector(context.Background(), "database")
		if err != nil {
			log.Fatalf("failed to initialize DB: %v", err)
			panic(err)
		}
		instance = sql.OpenDB(con)
	}

	return instance
}

func GetInstance() *sql.DB {
	getInstance()
	return instance
}
