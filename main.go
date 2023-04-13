package main

import (
	"awesomeProject/MarketData/Binance"
	"fmt"
)

func main() {

	binance := Binance.Binance{}
	err := binance.Init()
	if err != nil {
		fmt.Println(err.Error())
	}

	price, err := binance.GetSymbolPrice("BTCUSDT")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Btc price is", price)

	candleSticks, err := binance.GetKlines("BTCUSDT", "15m")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(candleSticks[0].High)
}
