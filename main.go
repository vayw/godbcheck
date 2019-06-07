package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func fail() {
	fmt.Println(0)
	os.Exit(1)
}

// VERSION from git, determined on build
var version = "undefined"

func main() {
	var (
		dbtype   = flag.String("t", "postgresql", "db type")
		host     = flag.String("h", "", "db hostname")
		user     = flag.String("u", "", "db user")
		password = flag.String("p", "", "db password")
		name     = flag.String("n", "", "db name")
		port     = flag.Int("port", -1, "port number")
		Err      error
		DB       *sql.DB
		ver      = flag.Bool("version", false, "print program version and exit")
	)
	flag.Parse()

	if *ver {
		fmt.Println(version)
		os.Exit(0)
	}

	switch *dbtype {
	case "mysql":
		var mport int
		if *port == -1 {
			mport = 3306
		} else {
			mport = *port
		}
		connStr := fmt.Sprintf("%s:%s@%s:%d/%s", *user, *password, *host, mport, *name)
		DB, Err = sql.Open("mysql", connStr)
	case "postgresql":
		var mport int
		if *port == -1 {
			mport = 5432
		} else {
			mport = *port
		}
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", *user, *password, *host, mport, *name)
		DB, Err = sql.Open("postgres", connStr)
	}
	if Err != nil {
		fail()
	}
	defer DB.Close()

	Err = DB.Ping()
	if Err != nil {
		fail()
	}
	fmt.Println(1)
}
