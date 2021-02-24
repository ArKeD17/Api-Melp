package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// Driver para postgresql
	_ "github.com/lib/pq"
	"github.com/mvochoa/logger"
)

// ConnectRoot genera una conexión a la base de datos con un usuario root
func ConnectRoot() *sql.DB {
	password := os.Getenv("POSTGRES_PASSWORD")
	return connect(os.Getenv("POSTGRES_USER"), string(password))
}

func connect(user, password string) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		user,
		password,
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("CONNECT DATABASE", err)
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.Error("PING DATABASE", err)
		log.Fatal(err)
	}

	return db
}
