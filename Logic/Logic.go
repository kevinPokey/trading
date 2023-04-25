package Logic

import "trading/Models"

type Logic interface {
	Init(candleSticks []Models.CandleStick, symbol string)

	GetRSI(timeFrame int) (err error)
}
