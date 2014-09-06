package main

import (
	"encoding/json"
	"fmt"
	"time"

	ldbpg "github.com/aranasaurus/leveldb-playground"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	dbPath := "../db"

	for {
		db, err := leveldb.OpenFile(dbPath, nil)
		if err != nil {
			fmt.Println(err)
			continue
		}

		itemCount := 0
		batch := new(leveldb.Batch)
		iter := db.NewIterator(nil, nil)
		for iter.Next() {

			hiJson, err := db.Get(iter.Key(), nil)
			if err != nil {
				fmt.Printf("Error reading item %v: %v\n", string(iter.Key()), err)
				continue
			}

			var hi ldbpg.HistoryItem
			if err := json.Unmarshal(hiJson, &hi); err != nil {
				fmt.Printf("Error unmarshalling item %v: %v\n", string(iter.Key()), err)
				fmt.Println(string(hiJson))
				continue
			}

			fmt.Println("Marking for deletion:", hi.UserId)
			batch.Delete(iter.Key())
			itemCount++
		}
		iter.Release()
		if err := iter.Error(); err != nil {
			fmt.Println("Error iterating db:", err)
		}

		if err := db.Write(batch, nil); err != nil {
			fmt.Println("Error deleting items:", err)
		} else {
			var s = "s"
			if itemCount == 1 {
				s = ""
			}
			if itemCount > 0 {
				fmt.Printf("Deleted %d item%v.\n", itemCount, s)
			}
		}

		db.Close()

		time.Sleep(2 * time.Second)
	}
}
