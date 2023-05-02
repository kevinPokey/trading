package Talib

import (
	"github.com/markcheno/go-talib"
)

func (tq *TaQuote) GetRSI(timeFrame int) []float64 {
	return talib.Rsi(tq.close, timeFrame)
}
