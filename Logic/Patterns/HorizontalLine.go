package Patterns

import (
	"math"
	"trading/Models"
)

func getAveragePrice(candles []Models.CandleStick) float64 {

	avg := 0.0
	for _, candle := range candles {
		avg = avg + (candle.Open+candle.Close)/2
	}

	return avg / float64(len(candles))
}

func findBottomHorizontalLine(candles []Models.CandleStick) []float64 {

	lowestPrice := math.MaxFloat64
	avg := getAveragePrice(candles)
	allowedDiff := avg * 0.02

	var downCandles []Models.CandleStick

	for _, candle := range candles {
		if !CandleDirection(&candle) {
			downCandles = append(downCandles, candle)
		}
	}

}
