package postgres

import (
	"context"
	"testing"

	"github.com/abeltay/go-template/models"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestAddUser(t *testing.T) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	// Given:
	expectName := "Jane Doe"
	svc := newFakeService(tx)

	// When:
	out, err := svc.AddUser(ctx, expectName)
	assert.NoError(t, err)

	// Then:
	user, err := models.FindUser(ctx, tx, out)
	assert.NoError(t, err)
	assert.Equal(t, expectName, user.FullName)
}

func TestUserFullName(t *testing.T) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	// Given:
	id := 100
	expectName := "John Doe"
	user := models.User{ID: id, FullName: expectName}
	if err = user.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Errorf("error setting up mocks, %s", err)
	}

	svc := newFakeService(tx)

	// When:
	out, err := svc.UserFullName(ctx, id)

	// Then:
	assert.NoError(t, err)
	assert.Equal(t, expectName, out)
}
