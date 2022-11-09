package main

import (
	"fmt"
	"go_docker/database/driver"
	"go_docker/database/users"
	"time"
)

func main() {
	fmt.Println("Hello golang from docker with air!", "and you")

	db := driver.Connect()

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println(db)
	users.All(db)
}
