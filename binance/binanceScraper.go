package binance

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func AllUSDTKlineScraper(binanceClient *binance.Client) {
	fmt.Println("*********************************************")
	fmt.Println("Binance All USDT Kline Scraper Started")
	fmt.Println("*********************************************")

	symbolList, _ := GetAllUSDTSymbolList(binanceClient)

	fmt.Println(symbolList)
	fmt.Println("All USDT Symbols:", len(symbolList))

	klineSymbol1mMapList := map[string]string{}
	klineSymbol3mMapList := map[string]string{}
	klineSymbol5mMapList := map[string]string{}
	klineSymbol15mMapList := map[string]string{}
	klineSymbol30mMapList := map[string]string{}
	klineSymbol1hMapList := map[string]string{}
	klineSymbol4hMapList := map[string]string{}
	klineSymbol1dMapList := map[string]string{}

	for i := range symbolList {
		klineSymbol1mMapList[symbolList[i]] = "1m"
		klineSymbol3mMapList[symbolList[i]] = "3m"
		klineSymbol5mMapList[symbolList[i]] = "5m"
		klineSymbol15mMapList[symbolList[i]] = "15m"
		klineSymbol30mMapList[symbolList[i]] = "30m"
		klineSymbol1hMapList[symbolList[i]] = "1h"
		klineSymbol4hMapList[symbolList[i]] = "4h"
		klineSymbol1dMapList[symbolList[i]] = "1d"
	}

	WsCombinedKlineStreamDataCollector(klineSymbol1mMapList)
}

func AllKlineScraper(binanceClient *binance.Client) {
	fmt.Println("*********************************************")
	fmt.Println("Binance All Kline Scraper Scraper Started")
	fmt.Println("*********************************************")

	symbolList, _ := GetAllSymbolList(binanceClient)

	fmt.Println(symbolList)
	fmt.Println("All Symbols:", len(symbolList))

	klineSymbol1mMapList := map[string]string{}
	klineSymbol3mMapList := map[string]string{}
	klineSymbol5mMapList := map[string]string{}
	klineSymbol15mMapList := map[string]string{}
	klineSymbol30mMapList := map[string]string{}
	klineSymbol1hMapList := map[string]string{}
	klineSymbol4hMapList := map[string]string{}
	klineSymbol1dMapList := map[string]string{}

	for i := range symbolList {
		klineSymbol1mMapList[symbolList[i]] = "1m"
		klineSymbol3mMapList[symbolList[i]] = "3m"
		klineSymbol5mMapList[symbolList[i]] = "5m"
		klineSymbol15mMapList[symbolList[i]] = "15m"
		klineSymbol30mMapList[symbolList[i]] = "30m"
		klineSymbol1hMapList[symbolList[i]] = "1h"
		klineSymbol4hMapList[symbolList[i]] = "4h"
		klineSymbol1dMapList[symbolList[i]] = "1d"
	}

	WsCombinedKlineStreamDataCollector(klineSymbol1mMapList)
}
