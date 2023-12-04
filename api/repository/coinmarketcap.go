package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"realTimeCryptoPrice/api/types"
)

func CoinMarketCapAPI(coin string) (types.CoinType, error) {
	url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?slug=%s&convert_id=2783", coin)
	apiKey := os.Getenv("API_KEY")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return types.CoinType{}, err
	}
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return types.CoinType{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return types.CoinType{}, err
	}

	var cryptoCoin types.CoinType
	err = json.Unmarshal(body, &cryptoCoin)
	if err != nil {
		return types.CoinType{}, err
	}

	return cryptoCoin, nil
}
