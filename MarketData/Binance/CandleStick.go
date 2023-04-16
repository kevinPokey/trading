package Binance

import (
	"context"
	"github.com/adshao/go-binance/v2"
	"trading/Models"
)

func MapToCandleSticks(klines []*binance.Kline) (candleSticks []Models.CandleStick) {
	//In Go, slices and maps are already reference types, so there is no need to return a pointer to them

	//make an array for candleSticks
	candleSticks = make([]Models.CandleStick, len(klines))

	//map klines to candlesticks
	for i, kline := range klines {
		candleStick := Models.CandleStick{
			OpenTime:  kline.OpenTime,
			CloseTime: kline.CloseTime,
			Open:      kline.Open,
			Close:     kline.Close,
			High:      kline.High,
			Low:       kline.Low,
			Volume:    kline.Volume,
		}
		candleSticks[i] = candleStick
	}
	return
}

func (b *Binance) GetCandlesticks(symbol string, interval string) (candleSticks []Models.CandleStick, err error) {

	//Get klines from binance
	klines, err := b.client.NewKlinesService().Symbol(symbol).Interval(interval).Do(context.Background())
	if err != nil {
		return
	}

	//Map klines to our own struct for compatibility
	candleSticks = MapToCandleSticks(klines)
	return
}
