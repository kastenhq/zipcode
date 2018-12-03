package zipcode

import (
	"database/sql"
	"fmt"
	"os"

	// Register the Postgres driver
	_ "github.com/lib/pq"
)

const (
	pgHostEnv     = "PG_HOST"
	pgPortEnv     = "PG_PORT"
	pgDBEnv       = "PG_DBNAME"
	pgUserEnv     = "PG_USER"
	pgPasswordEnv = "PG_PASSWORD"
	pgSSLEnv      = "PG_SSL"
)

func newPostgreSQLDB() (*sql.DB, error) {
	cfg, err := parseEnv()
	if err != nil {
		return nil, err

	}

	// Initialize connection object.
	db, err := sql.Open("postgres", cfg.String())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully created connection to database")
	return db, nil
}

type dbConfig struct {
	host     string
	port     string
	dbName   string
	user     string
	password string
}

func parseEnv() (*dbConfig, error) {
	host, ok := os.LookupEnv(pgHostEnv)
	if !ok {
		return nil, fmt.Errorf("%s environment variable not set", pgHostEnv)
	}
	port, ok := os.LookupEnv(pgPortEnv)
	if !ok {
		return nil, fmt.Errorf("%s environment variable not set", pgPortEnv)
	}
	dbName, ok := os.LookupEnv(pgDBEnv)
	if !ok {
		return nil, fmt.Errorf("%s environment variable not set", pgDBEnv)
	}
	user, ok := os.LookupEnv(pgUserEnv)
	if !ok {
		return nil, fmt.Errorf("%s environment variable not set", pgUserEnv)
	}
	password, ok := os.LookupEnv(pgPasswordEnv)
	if !ok {
		return nil, fmt.Errorf("%s environment variable not set", pgPasswordEnv)
	}
	return &dbConfig{host: host, port: port, dbName: dbName, user: user, password: password}, nil
}

func (c *dbConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.host, c.port, c.user, c.password, c.dbName)
}
