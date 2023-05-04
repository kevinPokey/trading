package Binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
)

type Binance struct {
	client *binance.Client
}

func (b *Binance) Init(apiKey string, secretKey string) (err error) {
	//Use testnet
	//binance.UseTestnet = true

	//Initialize client and test connection
	b.client = binance.NewClient(apiKey, secretKey)
	err = b.client.NewPingService().Do(context.Background())
	if err != nil {
		fmt.Println("Could not establish a connection with binance")
		return
	}

	return

}
