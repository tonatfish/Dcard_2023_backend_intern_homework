package model

type Page struct {
	Articles    []string `json:"articles"`
	NextPageKey string   `json:"nextPageKey"`
}
