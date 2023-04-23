package coinmarket

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Crypto struct {
	Coin struct {
		Usd float32 `json:"usd"`
	} `json:"ethereum"`
}

var data Crypto

func GetCurrentPrice(s string) float32 {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://"+os.Getenv("BASE_URL")+"/api/v3/simple/price", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := url.Values{}
	q.Add("ids", s)
	q.Add("vs_currencies", "USD")

	req.Header.Set("accept", "application/json")
	// req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("COIN_MARKET_API"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	// // body, err := ioutil.ReadAll(resp.Body)
	// fmt.Printf("status code: %s\n", resp.Status)
	// fmt.Println(resp.Body)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		// fmt.Println(bodyString)
		json.Unmarshal([]byte(bodyString), &data)

		// log.Info(bodyString)
	}
	return data.Coin.Usd

}
