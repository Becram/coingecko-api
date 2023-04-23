package main

import (
	"fmt"
	"os"
	"secretmanager2env/coinmarket"
	"secretmanager2env/email"
)

func main() {

	// crypto := flag.String("s", "changeme", "Secret To Fetch")
	// flag.Parse()
	// if *crypto == "changeme" {
	// 	fmt.Println("You must specify a crypto name to fetch")
	// 	return
	// }

	// var data map[string]string

	// json.Unmarshal([]byte(coinmarket.GetCurrentPrice()), &data)
	fmt.Println("Price for Ethereum:", coinmarket.GetCurrentPrice("ethereum"))

	data := &email.Email{From: os.Getenv("FROM_EMAIL"), To: os.Getenv("TO_EMAIL"), Content: "Price for Ethereum:" + fmt.Sprintf("%f", coinmarket.GetCurrentPrice("ethereum")) + "$", Subject: "Price Update For Ethereum", Api: os.Getenv("SENDGRID_API_KEY")}
	email.SendEmail(*data)

}
