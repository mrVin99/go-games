package database

import (
	"log"
	"os"
)

func migrate() {
	if tableExists("player") {
		return
	}

	file, err := os.ReadFile("schema.sql")
	if err != nil {
		panic(err)
	}

	if err = Conn().Exec(string(file)); err != nil {
		return
	}

	log.Println("Database migration concluded.")
}

func tableExists(tableName string) bool {
	var exists bool
	query := `SELECT EXISTS (
                    SELECT 1
                    FROM   information_schema.tables 
                    WHERE  table_schema = 'public'
                    AND    table_name = $1
            )`

	if err := Conn().
		QueryRow(query, tableName).
		Scan(&exists); err != nil {
		panic(err)
	}

	return exists
}
