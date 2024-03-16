package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Data struct {
	Error  []interface{}          `json:"error"`
	Prices map[string]interface{} `json:"result"`
	Last   int                    `json:"last"`
}

// Remove this if you dont need it, its used to get the entries to use them in other things ..
func getEntries(Data Data) ([]interface{}, error) {
	for _, entries := range Data.Prices {
		return entries.([]interface{}), nil
	}
	return nil, fmt.Errorf("entries not found")
}

func main() {
	pair, interval := "ETHUSD", 240
	url := fmt.Sprintf("https://api.kraken.com/0/public/OHLC?pair=%s&interval=%d", pair, interval)

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	var Data Data
	if err := json.NewDecoder(response.Body).Decode(&Data); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	// Printing parsed data
	fmt.Println("Last:", Data.Last)
	fmt.Println("Error:", Data.Error)
	for pair, entries := range Data.Prices {
		fmt.Printf("Pair: %s, Interval: %d\n", pair, interval)
		fmt.Println("Entries:")
		for _, entry := range entries.([]interface{}) {
			entrySlice := entry.([]interface{})
			timestamp := entrySlice[0].(float64)
			open := entrySlice[1].(string)
			high := entrySlice[2].(string)
			close := entrySlice[3].(string)
			low := entrySlice[4].(string)
			vwap := entrySlice[5].(string)
			volume := entrySlice[6].(string)
			count := entrySlice[7].(float64)

			fmt.Printf("  Timestamp: %s, Open: %s, High: %s, Close: %s, Low: %s, VWAP: %s, Volume: %s, Count: %.0f\n",
				strconv.FormatFloat(timestamp, 'f', 0, 64), open, high, close, low, vwap, volume, count)
		}
		break
	}

}
