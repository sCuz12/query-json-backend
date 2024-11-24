package data

import (
	"database/sql"
	"fmt"
	"os"

	db2 "github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mysql"
	"github.com/upper/db/v4/adapter/postgresql"
)

var db *sql.DB
var upper db2.Session

// Models is the wrapper for all database models
type Models struct {
	// any models inserted here (and in the New function)
	// are easily accessible throughout the entire application
	Users User
	Tokens Token
	AppStat AppStat
}

// New initializes the models package for use
func New(databasePool *sql.DB) Models {
	db = databasePool
	var err error
	fmt.Println(databasePool);
	switch os.Getenv("DATABASE_TYPE") {
	case "mysql", "mariadb":
		upper, _ = mysql.New(databasePool)
	case "postgres", "postgresql":
		upper, err  = postgresql.New(databasePool)
		if err != nil {
			fmt.Printf("Error initializing PostgreSQL session: %v\n", err)
			os.Exit(1)
		}
	default:
		// do nothing
	}

	return Models{}
}

// getInsertID returns the integer value of a newly inserted id (using upper)
func getInsertID(i db2.ID) int {
	idType := fmt.Sprintf("%T", i)
	if idType == "int64" {
		return int(i.(int64))
	}

	return i.(int)
}
