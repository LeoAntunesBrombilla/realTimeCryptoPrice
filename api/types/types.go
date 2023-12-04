package types

import "time"

type CoinType struct {
	Data struct {
		Num1 struct {
			ID          int       `json:"id"`
			Name        string    `json:"name"`
			Symbol      string    `json:"symbol"`
			Slug        string    `json:"slug"`
			IsActive    int       `json:"is_active"`
			LastUpdated time.Time `json:"last_updated"`
			Quote       struct {
				Num2783 struct {
					Price            float64   `json:"price"`
					VolumeChange24H  float64   `json:"volume_change_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					PercentChange30D float64   `json:"percent_change_30d"`
					PercentChange60D float64   `json:"percent_change_60d"`
					PercentChange90D float64   `json:"percent_change_90d"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"2783"`
			} `json:"quote"`
		} `json:"1"`
	} `json:"data"`
}
