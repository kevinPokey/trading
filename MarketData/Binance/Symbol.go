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

func (b *Binance) GetSymbols() (symbols []string, err error) {
	symbolTickers, err := b.client.NewListPricesService().Do(context.Background())
	if err != nil {
		return
	}

	//Convert array of struct to array of strings
	symbols = make([]string, 0, len(symbolTickers))
	for _, s := range symbolTickers {
		symbols = append(symbols, s.Symbol)
	}
	return
}
