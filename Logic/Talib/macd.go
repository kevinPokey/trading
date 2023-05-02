package Talib

import (
	"github.com/markcheno/go-talib"
)

// TRADING VIEW LINES
// result1 BLUE
// result2 ORANGE
// result3 MACD

func (tq *TaQuote) GetMACD(fastPeriod int, slowPeriod int, signalPeriod int) ([]float64, []float64, []float64) {
	return talib.Macd(tq.close, fastPeriod, slowPeriod, signalPeriod)
}
