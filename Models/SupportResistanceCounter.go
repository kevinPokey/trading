package Models

type SupportResistanceCounter struct {
	Price      float64
	Occurrence int
}

func (spc *SupportResistanceCounter) IsWithinPriceRange(price float64, allowedError float64) bool {
	if spc.Price-price < price || spc.Price+price > price {
		return true
	}
	return false
}
