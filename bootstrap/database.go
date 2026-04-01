package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"database/sql"
	_ "github.com/lib/pq"
)

func NewPostgresDatabase(env *Env) *sql.DB {

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName
	dbDriver := env.DBDriver

	postgresdbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open(dbDriver, postgresdbURI)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		log.Fatal(err)
	}
	return db
}

func ClosePostgresDBConnection(db *sql.DB) {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to PostgresDB closed.")
}
