package shortener

import (
	"github.com/dongri/go-shortener/bitly"
	"github.com/dongri/go-shortener/googl"
)

// NewBitly ...
func NewBitly(accessToken string) *bitly.Bitly {
	return bitly.New(accessToken)
}

// NewGoogl ...
func NewGoogl(apiKey string) *googl.Googl {
	return googl.New(apiKey)
}
