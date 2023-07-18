package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlx.DB
}

func New(config Config) (*Database, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database,
	)
	client, err := sqlx.Open("postgres", connInfo)
	if err != nil {
		return nil, err
	}

	err = client.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{client}, nil
}
