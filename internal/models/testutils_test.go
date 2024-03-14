package models

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newTestDBPool(t *testing.T) *pgxpool.Pool {
	dsn := "host=localhost port=5432 user=test_web password=pass dbname=test_snippetbox"
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	// Read the setup SQL script from file and execute the statements.
	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = dbpool.Exec(context.Background(), string(script))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbpool.Exec(context.Background(), string(script))
		if err != nil {
			t.Fatal(err)
		}
		dbpool.Close()
	})
	// Return the database connection pool.
	return dbpool
}
