package Binance

import (
	"context"
)

func (b *Binance) GetSymbolPrice(symbol string) (price string, err error) {

	ticker, err := b.client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		return
	}
	price = ticker[0].Price
	return
}
