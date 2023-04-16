package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"trading/Controllers"
	"trading/MarketData/Binance"
	"trading/Routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Could not load .env file")
	}

	apiKey := os.Getenv("API_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	//Initialize binance
	binance := &Binance.Binance{}
	err = binance.Init(apiKey, secretKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Initialize Controller
	controller := Controllers.NewController(binance)

	//Initialize Fiber
	app := fiber.New()

	router := Routes.CreateRouter(app, controller)
	router.RegisterRoutes()

	app.Listen(":3000")
}
