package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5435"
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {

	var db *sql.DB

	for i := 0; i < 5; i++ {
		var err error
		db, err = connectPostgres()
		if err != nil {
			fmt.Printf("Error connecting to database: %s\n", err)
		}
		if db != nil {
			fmt.Println("Connected...")
			break
		}
		time.Sleep(500 * time.Millisecond)

	}

	tick := time.Tick(100 * time.Millisecond)
	done := time.After(3 * time.Second)

	for {
		select {
		case <-tick:
			var id int
			query :=   `INSERT INTO person (name, surname, number)
						VALUES ($1, $2, $3)
						RETURNING id`

			err:=db.QueryRow(query, "John", "Doe", "123").Scan(&id)
			if err != nil {
				fmt.Printf("Error inserting record: %s\n", err)
			}else{
				fmt.Printf("%d. record inserted\n", id)
			}
			
			
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func connectPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
