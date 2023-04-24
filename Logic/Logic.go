package Logic

import "trading/Models"

type Logic interface {
	Init(candleSticks []Models.CandleStick, symbol string) (err error)

	GetRSI() (err error)
}
