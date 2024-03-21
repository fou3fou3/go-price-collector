package main

import (
	"fmt"

	krakenapi "github.com/beldur/kraken-go-api-client"
)

func main() {
	// Interval should be a string .
	pair, interval := "ETH/USD", "240"
	api := krakenapi.New("", "")

	resp, err := api.OHLCWithInterval(pair, interval)

	if err != nil {
		println("Error getting ohlc data:")
	}

	for _, ohlc := range resp.OHLC {
		fmt.Printf("\n\nTime: %v\n", ohlc.Time)
		fmt.Printf("Open: %f\n", ohlc.Open)
		fmt.Printf("High: %f\n", ohlc.High)
		fmt.Printf("Low: %f\n", ohlc.Low)
		fmt.Printf("Close: %f\n", ohlc.Close)

	}

}
