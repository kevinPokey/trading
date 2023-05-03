package Patterns

import "trading/Models"

func CandleDirection(candle *Models.CandleStick) bool {
	if candle.Open < candle.Close {
		return true
	}
	return false
}

func TrendDirection(cadles []Models.CandleStick) {

}
