package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbUser         = "root"
	dbPassword     = ""
	database       = "simple-bank"
	dbRootPassword = ""
)

const (
	dbDriver = "mysql"
	dbSource = "root@localhost/simple-bank"
)

func TestMain(m *testing.M) {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, "localhost", "3306", database)
	fmt.Println(dbUri)
	conn, err := sql.Open(dbDriver, dbUri)

	if err != nil {
		log.Fatal("fail to connect")
	}
	testQueries = New(conn)
	os.Exit(m.Run())

}
