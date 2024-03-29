package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	// "go_docker/database/driver"
	// "go_docker/database/users"
	// "time"
)

func main() {
	fmt.Println("Hello golang from docker with air!", "and you")
	// send request to external api
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	fmt.Println(resp)
	/*
	 * run as web server
	 * listen request
	 * access to http://localhost:8000
	 */
	http.HandleFunc("/", handler)
	// for Graphql playground, modify port
	// Note: avoid conflict to server.go's Port setting
	http.ListenAndServe(":8000", nil)
	/**
	 * db connection
	 */
	// db := driver.Connect()
	// See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)
	// fmt.Println(db)
	// users.All(db)
}

func handler(w http.ResponseWriter, r *http.Request) {
	origin := r.Referer()
	fmt.Println("Origin:", origin)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	message := "Hello, World from golang! "
	message += time.Now().Format("2006-01-02 15:04:05")
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)
}
