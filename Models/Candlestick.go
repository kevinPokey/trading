package Models

type CandleStick struct {
	OpenTime  int64
	CloseTime int64
	Open      float64
	Close     float64
	High      float64
	Low       float64
	Volume    float64
}
