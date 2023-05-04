package Binance

import (
	"context"
	"github.com/adshao/go-binance/v2"
	"strconv"
)

func (b *Binance) GetBalance() (balances []binance.Balance, err error) {
	res, err := b.client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return
	}
	balances = res.Balances
	return
}

func (b *Binance) GetOrders() ([]*binance.Order, error) {
	return b.client.NewListOpenOrdersService().Do(context.Background())
}

func (b *Binance) PlaceOrder(symbol string, usdt float64) (*binance.CreateOrderResponse, error) {
	return b.client.NewCreateOrderService().Symbol(symbol).Side(binance.SideTypeBuy).
		Type(binance.OrderTypeMarket).QuoteOrderQty(strconv.FormatFloat(usdt, 'f', -1, 64)).Do(context.Background())
}

func (b *Binance) CancelOrder(id int) (*binance.CancelOrderResponse, error) {
	return b.client.NewCancelOrderService().OrderID(int64(id)).Do(context.Background())
}
