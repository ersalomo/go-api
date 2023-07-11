package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"simple-bank/utils"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	dbDriver, dbUri := utils.GetDbUrl()
	fmt.Println(dbUri)
	conn, err := sql.Open(dbDriver, dbUri)

	if err != nil {
		log.Fatal("fail to connect")
	}
	testQueries = New(conn)
	os.Exit(m.Run())

}
