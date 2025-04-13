package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

var DB *sqlx.DB

func InitDB() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbName)

	for i := 0; i < 10; i++ {
		DB, err = sqlx.Connect("mysql", dataSource)
		if err == nil {
			break
		}
		fmt.Println("Waiting for DB to be ready...", err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	createTables()
}

func createTables() {
	createTodosTable := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER AUTO_INCREMENT PRIMARY KEY ,
        name NVARCHAR(50) NOT NULL,
        description NVARCHAR(255) NOT NULL,
        startDate DATETIME NOT NULL,
        endDate DATETIME,
        completed BOOLEAN NOT NULL,
        userId INT NOT NULL REFERENCES users(id)
    )
    `

	_, err := DB.Exec(createTodosTable)

	if err != nil {
		panic(err)
		return
	}

	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
      id INTEGER AUTO_INCREMENT PRIMARY KEY,
      email VARCHAR(255) NOT NULL,
      password VARCHAR(255) NOT NULL
    )   
    `

	_, err = DB.Exec(createUsersTable)

	if err != nil {
		panic(err)
		return
	}

}
