package main

import (
	"bufio"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"trading/Controllers"
	"trading/MarketData/Binance"
	"trading/Routes"
)

func readConstFile(fileName string) (result []string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := scanner.Text()
		result = append(result, r)
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

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
	constIndicators, err := readConstFile("Consts/indicators.txt")
	if err != nil {
		return
	}
	constIntervals, err := readConstFile("Consts/intervals.txt")
	if err != nil {
		return
	}
	constTickers, err := readConstFile("Consts/tickers.txt")
	if err != nil {
		return
	}
	controller := Controllers.NewController(binance, constIndicators, constIntervals, constTickers)

	//Initialize Fiber
	app := fiber.New()

	router := Routes.CreateRouter(app, controller)
	router.RegisterRoutes()

	fmt.Println(strconv.FormatFloat(104.4, 'f', 8, 64))

	app.Listen(":3000")

}
