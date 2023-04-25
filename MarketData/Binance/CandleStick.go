package Binance

import (
	"context"
	"github.com/adshao/go-binance/v2"
	"strconv"
	"trading/Models"
)

func mapToCandleSticks(klines []*binance.Kline) ([]Models.CandleStick, error) {
	//In Go, slices and maps are already reference types, so there is no need to return a pointer to them

	//make an array for candleSticks
	candleSticks := make([]Models.CandleStick, len(klines))

	//map klines to candlesticks
	for i, kline := range klines {

		openVal, err := strconv.ParseFloat(kline.Open, 64)
		if err != nil {
			return candleSticks, err
		}
		closeVal, err := strconv.ParseFloat(kline.Close, 64)
		if err != nil {
			return candleSticks, err
		}
		highVal, err := strconv.ParseFloat(kline.High, 64)
		if err != nil {
			return candleSticks, err
		}
		lowVal, err := strconv.ParseFloat(kline.Low, 64)
		if err != nil {
			return candleSticks, err
		}
		volume, err := strconv.ParseFloat(kline.Volume, 64)
		if err != nil {
			return candleSticks, err
		}

		candleStick := Models.CandleStick{
			OpenTime:  kline.OpenTime,
			CloseTime: kline.CloseTime,
			Open:      openVal,
			Close:     closeVal,
			High:      highVal,
			Low:       lowVal,
			Volume:    volume,
		}
		candleSticks[i] = candleStick
	}
	return candleSticks, nil
}

func (b *Binance) GetCandlesticks(symbol string, interval string) (candleSticks []Models.CandleStick, err error) {

	//Get klines from binance
	klines, err := b.client.NewKlinesService().Symbol(symbol).Interval(interval).Do(context.Background())
	if err != nil {
		return
	}

	//Map klines to our own struct for compatibility
	candleSticks, err = mapToCandleSticks(klines)
	return
}
