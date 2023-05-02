package Logic

import "trading/Models"

type Logic interface {
	Init(candleSticks []Models.CandleStick, symbol string)

	GetRSI(timeFrame int) []float64
	GetMACD(fastPeriod int, slowPeriod int, signalPeriod int) ([]float64, []float64, []float64)
}
