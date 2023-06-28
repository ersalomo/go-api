package db

import (
	"database/sql"
	_ "github.com/lib/pg"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5431/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
