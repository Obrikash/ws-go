package user

import (
	"context"
	"database/sql"
)

type DBTX interface {
    ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
    PrepareContext(context.Context, string) (*sql.Stmt, error)
    QueryContext(context.Context, string, ...any) (*sql.Rows, error)
    QueryRowContext(context.Context, string, ...any) *sql.Row
}

type repository struct {
    db DBTX
}

func NewRepository(db DBTX) Repository {
    return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
    query := `INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id`

    var lastInsertedId int64

    err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertedId)
    if err != nil {
        return &User{}, err
    }

    user.ID = lastInsertedId
    return user, nil
}
