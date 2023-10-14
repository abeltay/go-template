package postgres

import (
	"context"
	"fmt"

	"github.com/abeltay/go-template/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// AddUser adds a new user
func (f Service) AddUser(ctx context.Context, fullName string) (int, error) {
	tx, err := f.beginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("insert user, opening transaction: %w", err)
	}
	defer tx.Rollback()

	user := models.User{FullName: fullName}
	if err = user.Insert(ctx, tx, boil.Infer()); err != nil {
		return 0, fmt.Errorf("insert user: %w", err)
	}
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("tx commit: %w", err)
	}
	return user.ID, err
}

// UserFullName retrieves a user's FullName
func (f Service) UserFullName(ctx context.Context, id int) (string, error) {
	user, err := models.FindUser(ctx, f.db, id, models.UserColumns.FullName)
	if err != nil {
		return "", fmt.Errorf("retrieve user: %w", err)
	}
	return user.FullName, err
}
