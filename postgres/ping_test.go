package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	svc := newFakeService(tx)
	err = svc.Ping(ctx)
	assert.NoError(t, err)
}
