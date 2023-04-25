package Talib

import (
	"trading/Models"
)

type TaQuote struct {
	symbol string
	open   []float64
	high   []float64
	low    []float64
	close  []float64
	volume []float64
}

func (tq *TaQuote) Init(candleSticks []Models.CandleStick, symbol string) {

	tq.symbol = symbol

	for _, candle := range candleSticks {
		tq.low = append(tq.low, candle.Low)
		tq.high = append(tq.high, candle.High)
		tq.open = append(tq.open, candle.Open)
		tq.close = append(tq.close, candle.Close)
		tq.volume = append(tq.volume, candle.Volume)
	}
	return
}
