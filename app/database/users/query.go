package users

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	ID            int
	Name          string
	Email         string
	EmailVerified string
	CreatedAt     string
	UpdatedAt     string
}

func All(db *sql.DB) {
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
