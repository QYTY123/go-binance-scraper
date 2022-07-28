package main

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
	mybinance "github.com/shing1211/go-binance-scraper/binance"
	utils "github.com/shing1211/go-binance-scraper/utils"
)

func main() {
	// dummy  or empty api and secretkey will be enought for binance socket scraper
	var APIKEY = utils.GetOSEnv("BINANCE_APIKEY", "YOUR_API_KEY")
	var SECRETKEY = utils.GetOSEnv("BINANCE_SECRET_KEY", "YOUR_SECRET_KEY")

	binanceClient := binance.NewClient(APIKEY, SECRETKEY)

	if binanceClient == nil {
		fmt.Println("Connection to Binance failed")
		return
	} else {
		fmt.Println("Connected to Binance")
	}

	mybinance.AllUSDTKlineScraper(binanceClient)
	//mybinance.AllKlineScraper(binanceClient)
}
