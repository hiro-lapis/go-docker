package driver

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

func Connect() (db *sql.DB) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	c := mysql.Config{
		DBName:    "db",
		User:      "user",
		Passwd:    "pw",
		Addr:      "go-docker_db_1:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, _ = sql.Open("mysql", c.FormatDSN())
	return
}
