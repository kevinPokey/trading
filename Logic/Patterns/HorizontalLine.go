package Patterns

import (
	"trading/Models"
)

func getAveragePrice(candles []Models.CandleStick) float64 {

	avg := 0.0
	for _, candle := range candles {
		avg = avg + (candle.Open+candle.Close)/2
	}

	return avg / float64(len(candles))
}

func findBottomHorizontalLine(candles []Models.CandleStick) float64 {

	//avg := getAveragePrice(candles)
	//allowedDiff := avg * 0.03
	//lowestPrice := math.MaxFloat64

	//var line []Models.CandleStick

	//for _, candle := range candles {
	//	if !CandleDirection(&candle) && lowestPrice > candle.Close {

	//	}
	//}
	return 0.0
}
