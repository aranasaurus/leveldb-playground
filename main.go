package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

var dbPath = flag.String("dbPath", "./db", "Path to the db file")

type HistoryItem struct {
	Token                string     `json:"token"`
	UserId               string     `json:"userId"`
	DeviceId             string     `json:"deviceId"`
	AppId                string     `json:"appId"`
	TracksActivityIds    []string   `json:"tracksActivityIds"`
	GeofencesActivityIds []string   `json:"geofencesActivityIds"`
	Locations            []Location `json:"locations"`
}

type Location struct {
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	Accuracy    float64    `json:"accuracy"`
	ActivityIds []string   `json:"activityIds"`
	Geofences   []Geofence `json:"geofences"`
}

type Geofence struct {
	Message      string `json:"message"`
	ActivityId   string `json:"activityId"`
	FeatureLayer string `json:"featureLayer"`
	FeatureId    int    `json:"featureId"`
	Action       struct {
		Type string `json:"type"`
	} `json:"action"`
}

func init() {
	flag.Parse()
}

func main() {
	db, err := leveldb.OpenFile(*dbPath, nil)
	for err != nil {
		time.Sleep(100 * time.Millisecond)
		db, err = leveldb.OpenFile(*dbPath, nil)
		continue
	}
	defer db.Close()

	file, err := ioutil.ReadFile("./test.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	var items []HistoryItem
	if err = json.Unmarshal(file, &items); err != nil {
		panic(err)
	}

	for _, hi := range items {
		hiJson, _ := json.Marshal(hi)
		t := time.Now()
		if key, err := t.MarshalText(); err != nil {
			fmt.Println("Error making key:", err)
			continue
		} else {
			db.Put(key, hiJson, nil)
		}
	}
}
