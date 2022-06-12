package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Ad3bay0c/backend-tracing-go/application/trace"
	"github.com/Ad3bay0c/backend-tracing-go/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUsers(ctx context.Context) ([]models.User, error)
}

type db struct {
	dbConn *sqlx.DB
}

func NewDB() (DB, error) {
	DbConn, err := sqlx.Open("sqlite3", "./app.db")
	if err = DbConn.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Database Connected")
	return db{
		dbConn: DbConn,
	}, nil
}

func (repo db) CreateUser(ctx context.Context, user models.User) error {
	ctx, span := trace.NewSpan(ctx, "CreateUser DB Method", nil)
	defer span.End()

	_, err := repo.dbConn.
		ExecContext(ctx, "INSERT INTO users(id, name, email) VALUES(?, ?, ?)",
			user.ID, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (repo db) GetUsers(ctx context.Context) ([]models.User, error) {
	ctx, span := trace.NewSpan(ctx, "GetUser DB Method", nil)
	defer span.End()

	users := []models.User{}
	err := repo.dbConn.SelectContext(ctx, &users, "SELECT * FROM user")
	if err != nil {
		trace.AddSpanError(span, fmt.Errorf("internal server error: %v - line 56", err))
		trace.FailSpan(span, fmt.Sprintf("internal server error: %v - line 56", err))

		return nil, err
	}
	return users, err
}
