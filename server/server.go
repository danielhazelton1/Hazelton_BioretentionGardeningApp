/*
This file provides the entry point for the golang server.
Funcitons contained in this file configure the server to scan environment variables,
connect to the database, and configure function handlers for handling incoming GET/POST requests.
*/
package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db         *sql.DB
	dbPort     string
	serverPort string
	username   string
	password   string
	hostname   string
	dbname     string
)

func main() {
	scanEnvVariables()

	// CONNECT TO DATABASE
	fmt.Println("Connecting to database...")
	waitForDB()

	// SETUP HANDLERS
	fmt.Println("Configuring handlers...")
	configureHandlers()

	// START LISTENING
	fmt.Println("Server is started and listening:")
	startServer()
}

func startServer() {
	// open http ports
	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func configureHandlers() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer)) // give access to static files (ie. css)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/download", downloadPageHandler)
	http.HandleFunc("/downloadcsv", downloadFileHandler)
	http.HandleFunc("/", defaultHandler) // open index.html by default
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("default handler invoked")
	http.ServeFile(w, r, "dynamic/index.html")
}

func downloadPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("download page handler invoked")
	http.ServeFile(w, r, "dynamic/download.html")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("submit handler invoked")

	err := parseAndSubmit(r)
	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "dynamic/failure.html")
	} else {
		http.ServeFile(w, r, "dynamic/submit.html")
	}
}

// continuously attempts to connect to the database
func waitForDB() {
	db = nil
	for {
		err := dbConnect()
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
		fmt.Println("Couldn't reach database:", err)
		fmt.Println("Attempting to reconnect to database.")
	}
	fmt.Println("Successfully connected to database.")
}

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("download CSV invoked")

	err := downloadCSV(w, r)
	if err != nil {
		fmt.Println(err)
		http.ServeFile(w, r, "dynamic/failure.html")
	}
}

func dbConnect() (err error) {
	dbsource := username + ":" + password + "@tcp(" + hostname + ":" + dbPort + ")/" + dbname
	db, err = sql.Open("mysql", dbsource)
	if err != nil {
		return err
	}

	// check that the database is up and running
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

// reads in and stores environment variables from '.env' file
// unless the program was run with the '-local' flag
// ie. "go run . -local"
func scanEnvVariables() {
	hostname = "db"
	dbPort = "3306"

	// if -local tag is run, then set env vars
	if len(os.Args) > 1 && os.Args[1] == "-local" {
		os.Setenv("MYSQL_DATABASE", "gardenapp")
		os.Setenv("MYSQL_USER", "user1")
		os.Setenv("MYSQL_PASSWORD", "usbw")
		os.Setenv("SERVER_PORT", "80")
		hostname = "localhost"
	}

	// read in environment variables to global vars
	serverPort = string(os.Getenv("SERVER_PORT"))
	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	dbname = os.Getenv("MYSQL_DATABASE")
}
