package user

import (
	"database/sql"
	domain "godo/internal/domain/user"
)

type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (db *PostgresUserRepo) Add(user *domain.User) error {
	//_, err := db.db.Exec("INSERT INTO users values($1,$2,$3,$4)", user.Id, user.Name, user.Email, user.CreatedAt)
	return nil
}
