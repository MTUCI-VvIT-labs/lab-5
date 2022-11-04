package pg

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectToDB(url string) error {
	cfg, err := getDBConnectionConfig(url)
	if err != nil {
		return err
	}
	log.Println("successfully received DB config")

	cfg.MaxConns = 10

	err = createDbPool(cfg)
	if err != nil {
		return err
	}

	return nil
}

func getDBConnectionConfig(url string) (*pgxpool.Config, error) {
	if url == "" {
		return nil, errors.New("DATABASE_URL environment variable isn't exist")
	}
	poolConfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, errors.New("failed to parse connection string to pgxpool.Config")
	}
	return poolConfig, nil
}

func createDbPool(poolConfig *pgxpool.Config) error {
	var err error
	DB, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return errors.New("failed to connect to DB")
	}
	log.Println("successfully connected to the database")
	err = checkDBConnection()
	if err != nil {
		return errors.New("failed to ping server")
	}
	log.Println("successfully ping DB")
	return nil
}

func checkDBConnection() error {
	return DB.Ping(context.Background())
}
