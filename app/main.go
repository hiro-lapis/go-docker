package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type User struct {
	ID            int
	Name          string
	Email         string
	EmailVerified string
	CreatedAt     string
	UpdatedAt     string
}

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	fmt.Println("Hello golang from docker with air!", "and you")
	const DSN = "user:pw@tcp(go-docker_db_1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println(db)
	getRows(db)
}

func getRows(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, email, email_verified, created_at, updated_at  FROM users")
	if err != nil {
		log.Fatalf("getRows db.Query error err:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		u := &User{}
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.EmailVerified, &u.CreatedAt, &u.UpdatedAt); err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		fmt.Println(u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err:%v", err)
	}
}
