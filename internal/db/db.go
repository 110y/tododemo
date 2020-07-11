package db

import (
	"context"

	"github.com/110y/tododemo/internal/todo"
)

type DB interface {
	PutTODO(ctx context.Context, t *todo.TODO) error
}
