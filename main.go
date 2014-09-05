package main

import (
	"flag"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var dbPath = flag.String("dbPath", "db", "Path to the db file")

func init() {
	flag.Parse()
}

func main() {
	db, err := leveldb.OpenFile(*dbPath, nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	log.Printf("Using db file: %#v", *dbPath)
}
