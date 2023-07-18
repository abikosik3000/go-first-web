package healthchek

import (
	"fist-web-go/internal/db/postgres"
	"fmt"
)

type Repository struct {
	db *postgres.Database
}

func NewRepository(db *postgres.Database) *Repository {
	return &Repository{db: db}
}

func (repo Repository) Application() string {
	return "pong"
}

func (repo Repository) Database() error {
	var testOne []int
	err := repo.db.Select(&testOne, "SELECT 1")
	if err != nil {
		return fmt.Errorf("Database ping %w", err)
	}
	fmt.Println(testOne)
	return nil
}
