package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func ConnectDB() *sql.DB {
	dbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return db
}

func AddNewFile(filename string) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `
	insert into files(file_name)
	values ($1)`

	_, err := db.Query(sqlStatement, filename)
	if err != nil {
		return err
	}
	return nil
}

func GetFileNames() ([]string, error) {
	db := ConnectDB()
	defer db.Close()

	var names []string
	var name string

	sqlStatement := `
	select file_name from files`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		if err = rows.Scan(&name); err != nil {
			return []string{}, err
		}
		names = append(names, name)
	}

	return names, nil
}

func DeleteAllFiles() error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `
	truncate files`

	_, err := db.Query(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(filename string) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `
	delete from files
	where file_name = $1`

	_, err := db.Query(sqlStatement, filename)
	if err != nil {
		return err
	}
	return nil
}
