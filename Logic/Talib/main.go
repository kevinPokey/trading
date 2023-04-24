package Talib

import (
	"errors"
	"reflect"
	"trading/Models"
)

type TaQuote struct {
	Symbol string
	Open   []float64
	High   []float64
	Low    []float64
	Close  []float64
	Volume []float64
}

func (tq *TaQuote) Init(candleSticks []Models.CandleStick, symbol string) (err error) {

	tq.Symbol = symbol

	for _, candle := range candleSticks {
		v := reflect.ValueOf(candle)
		for i := 0; i < v.NumField(); i++ {
			val := v.Field(i).Float()
			switch i {
			case 0:
				tq.Open = append(tq.Open, val)
			case 1:
				tq.Close = append(tq.Close, val)
			case 2:
				tq.High = append(tq.High, val)
			case 3:
				tq.Low = append(tq.Low, val)
			case 4:
				tq.Volume = append(tq.Volume, val)
			default:
				err = errors.New("invalid field")
			}
		}
	}
	return
}
