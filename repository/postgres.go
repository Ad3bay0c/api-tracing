package repository

import (
	"log"

	"github.com/Ad3bay0c/backend-tracing-go/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	CreateUser(user models.User) error
}

type db struct {
	dbConn *sqlx.DB
}

func NewDB() (DB, error) {
	DbConn, err := sqlx.Open("sqlite3", "./app.db")
	if err = DbConn.Ping(); err != nil {
		return nil, err
	}
	log.Println("Database Connected")
	return db{
		dbConn: DbConn,
	}, nil
}

func (repo db) CreateUser(user models.User) error {
	_, err := repo.dbConn.
		Query("INSERT INTO users(id, name, email) VALUES(?, ?, ?)",
			user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}
