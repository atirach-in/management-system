package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnectDB() {
	DB_DSN := os.Getenv("DB_DSN")

	fmt.Print("\nDB_DSN: ", DB_DSN)

	db, err := sql.Open("mysql", DB_DSN)

	if err != nil {
		fmt.Print("\nConnect Failed")
		panic(err.Error())
	} else {
		fmt.Print("\nConnect Success")
	}
	defer db.Close()
}

func ConnectDB() *sql.DB {
	DB_DSN := os.Getenv("DB_DSN")
	db, err := sql.Open("mysql", DB_DSN)

	if err != nil {
		fmt.Print("\nConnect Failed")
		panic(err.Error())
	}

	return db
}
