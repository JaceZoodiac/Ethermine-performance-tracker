package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"time"
	"encoding/json"
)

type Data struct {
	Time int64
	ReportedHashrate int64
	Unpaid int64
}

type Stats struct {
	Data Data
}

func constructURL(address string) string {
	minerAddress := address
	url := fmt.Sprintf("https://api.ethermine.org/miner/:%s/currentStats", minerAddress)
	return url
}

func queryStats(address string) Stats {
	url := constructURL(address)

	res, getErr := http.Get(url)
	if getErr != nil {
		fmt.Println("Get failed!")
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println("Read failed!")
	}

	var currentStats Stats
	json.Unmarshal([]byte(string(body)), &currentStats)
	return currentStats
}

func main()  {
	minerAddress := "0x13b8fd56ce90f401e00f9e3e5a8d9c843739bfb3"

	currentStats := queryStats(minerAddress)

	fmt.Println("Time: ", currentStats.Data.Time)
	fmt.Println("Reported Hash: ", currentStats.Data.ReportedHashrate)
	fmt.Println("Unpaid: ", currentStats.Data.Unpaid)

	for {
		time.Sleep(15 * time.Minute)

		currentStats := queryStats(minerAddress)

		fmt.Println("Time: ", currentStats.Data.Time)
		fmt.Println("Reported Hash: ", currentStats.Data.ReportedHashrate)
		fmt.Println("Unpaid: ", currentStats.Data.Unpaid)
	}
}
