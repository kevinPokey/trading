package Binance

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/joho/godotenv"
	"os"
)

type Binance struct {
	client *binance.Client
}

func (b *Binance) Init() (err error) {
	//Use testnet
	binance.UseTestnet = true

	//Load env variables
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Println("Could not set env variables.")
		return
	}
	apiKey := os.Getenv("API-KEY")
	secretKey := os.Getenv("SECRET-KEY")

	//Initialize client and test connection
	b.client = binance.NewClient(apiKey, secretKey)
	err = b.client.NewPingService().Do(context.Background())
	if err != nil {
		fmt.Println("Could not establish a connection with binance")
		return
	}
	return

}
