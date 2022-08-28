package postgres

import "context"

// Ping sends a SELECT 1 to the database as a connectivity test
func (f Service) Ping(ctx context.Context) error {
	_, err := f.db.ExecContext(ctx, "SELECT 1")
	return err
}
