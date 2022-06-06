package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jerry-yue/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

// viper insteaded!!!
// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:P@ssw0rd@127.0.0.1:5432/simple_bank?sslmode=disable"
// )
// viper insteaded!!!

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Load config file failed: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
