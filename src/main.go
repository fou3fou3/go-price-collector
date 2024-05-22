package main

import (
	"fmt"
	"strconv"

	krakenapi "github.com/beldur/kraken-go-api-client"
)

func FetchPrices(pair string, interval int16) *krakenapi.OHLCResponse {
	sinterval := strconv.Itoa(int(interval))
	api := krakenapi.New("", "")

	resp, err := api.OHLCWithInterval(pair, sinterval)

	if err != nil {
		println("Error getting ohlc data: %s", err)
	}

	return resp
}

func main() {
	resp := FetchPrices("ETH/USD", 240)

	for _, ohlc := range resp.OHLC {
		fmt.Printf("\n\nTime: %v\n", ohlc.Time)
		fmt.Printf("Open: %f\n", ohlc.Open)
		fmt.Printf("High: %f\n", ohlc.High)
		fmt.Printf("Low: %f\n", ohlc.Low)
		fmt.Printf("Close: %f\n", ohlc.Close)
	}
}
