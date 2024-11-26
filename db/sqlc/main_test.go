package db

import (
	"context"
	"log"
	"os"
	"simple_bank/util"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	// Create a connection pool instead of a single connection
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}
	defer testDB.Close() // Close the pool when done

	testQueries = New(testDB) // Pass the pool to New

	os.Exit(m.Run())
}
