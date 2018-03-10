package bitly

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

const (
	baseURI     = "https://api-ssl.bitly.com"
	pathShorten = "/v3/shorten"
	pathExpand  = "/v3/expand"
)

// Bitly ...
type Bitly struct {
	AccessToken string
	Timeout     int
}

// New ...
func New(accessToken string, timeout ...int) *Bitly {
	b := new(Bitly)
	b.AccessToken = accessToken
	if len(timeout) > 0 {
		b.Timeout = timeout[0]
	} else {
		b.Timeout = 10
	}
	return b
}

// Short ...
func (b *Bitly) Short(longURL string) (string, error) {
	params := map[string][]string{
		"longUrl":      {longURL},
		"access_token": {b.AccessToken},
		"format":       {"txt"},
	}
	vals := url.Values(params)
	req, err := http.NewRequest("GET", baseURI, nil)
	if err != nil {
		return longURL, err
	}
	req.URL.Path = path.Join(req.URL.Path, pathShorten)
	req.URL.RawQuery = vals.Encode()
	httpClient := &http.Client{Timeout: time.Duration(b.Timeout) * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSpace(string(body)), nil
}

// Long ...
func (b *Bitly) Long(shortURL string) (string, error) {
	params := map[string][]string{
		"shortUrl":     {shortURL},
		"access_token": {b.AccessToken},
		"format":       {"txt"},
	}
	vals := url.Values(params)
	req, err := http.NewRequest("GET", baseURI, nil)
	if err != nil {
		return shortURL, err
	}
	req.URL.Path = path.Join(req.URL.Path, pathExpand)
	req.URL.RawQuery = vals.Encode()
	httpClient := &http.Client{Timeout: time.Duration(b.Timeout) * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		return shortURL, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return shortURL, err
	}
	return strings.TrimSpace(string(body)), nil
}
