package clacky

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	baseURI     = "https://clacky.org"
	pathShorten = "/api/shorten"
	pathExpand  = "/api/expand"
)

// Clacky ...
type Clacky struct {
	AccessToken string
	Timeout     int
}

// New ...
func New(accessToken string, timeout ...int) *Clacky {
	c := new(Clacky)
	c.AccessToken = accessToken
	if len(timeout) > 0 {
		c.Timeout = timeout[0]
	} else {
		c.Timeout = 10
	}
	return c
}

// Short ...
func (c *Clacky) Short(longURL string) (string, error) {
	params := map[string][]string{
		"long_url":     {longURL},
		"access_token": {c.AccessToken},
	}
	vals := url.Values(params)
	req, err := http.NewRequest("GET", baseURI, nil)
	if err != nil {
		return longURL, err
	}
	req.URL.Path = path.Join(req.URL.Path, pathShorten)
	req.URL.RawQuery = vals.Encode()
	httpClient := &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	type ResBody struct {
		StatusCode int    `json:"status_code"`
		ShortURL   string `json:"short_url"`
	}
	var resBody ResBody
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		fmt.Println(err)
	}
	return resBody.ShortURL, nil
}

// Long ...
func (c *Clacky) Long(shortURL string) (string, error) {
	params := map[string][]string{
		"short_url":    {shortURL},
		"access_token": {c.AccessToken},
	}
	vals := url.Values(params)
	req, err := http.NewRequest("GET", baseURI, nil)
	if err != nil {
		return shortURL, err
	}
	req.URL.Path = path.Join(req.URL.Path, pathExpand)
	req.URL.RawQuery = vals.Encode()
	httpClient := &http.Client{Timeout: time.Duration(c.Timeout) * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	type ResBody struct {
		StatusCode int    `json:"status_code"`
		LongURL    string `json:"long_url"`
	}
	var resBody ResBody
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		fmt.Println(err)
	}
	return resBody.LongURL, nil
}
