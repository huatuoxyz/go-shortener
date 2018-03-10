# go-shortener

Simple Bitly, Google URL Shortener API client for Go.

### Usage
```go
$ go get github.com/dongri/go-shortener
```

```go
package main

import (
	"fmt"
	shortener "github.com/dongri/go-shortener"
)

func main() {

	// bitly
	b := shortener.NewBitly("{ACCESS_TOKEN}")

	shortURL, err := b.Short("http://hoge.com")
	fmt.Println(shortURL) // http://bit.ly/2DahJMA

	longURL, err := b.Long(shortURL)
	fmt.Println(longURL) // http://hoge.com/

	// goo.gl
	g := shortener.NewGoogl("{GOOGLE_API_KEY}")

	shortURL, err = g.Short("http://hoge.com")
	fmt.Println(shortURL) // https://goo.gl/hQMZ6

	longURL, err = g.Long(shortURL)
	fmt.Println(longURL) // http://hoge.com/
}
```
