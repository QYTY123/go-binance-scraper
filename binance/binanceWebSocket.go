package binance

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

func WsCombinedAggTradeStreamsDataCollector(symbols []string) {
	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
		//controllers.AddAggTrade(event)
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsCombinedAggTradeServe(symbols, wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func WsCombinedKlineStreamDataCollector(symbolIntervalPair map[string]string) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		// Persist and print IsFinal data to db and screen only
		//if event.Kline.IsFinal {
		//controllers.AddKline(event)
		fmt.Println(event)
		//}
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsCombinedKlineServe(symbolIntervalPair, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}
