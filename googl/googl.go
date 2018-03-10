package googl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	baseURI       = "https://www.googleapis.com"
	pathShortener = "/urlshortener/v1/url"
)

// Googl ...
type Googl struct {
	APIKey  string
	Timeout int
}

// New ...
func New(apiKey string, timeout ...int) *Googl {
	g := new(Googl)
	g.APIKey = apiKey
	if len(timeout) > 0 {
		g.Timeout = timeout[0]
	} else {
		g.Timeout = 10
	}
	return g
}

// Short ...
func (g *Googl) Short(longURL string) (string, error) {
	jsonStr := `{"longUrl":"` + longURL + `"}`
	req, err := http.NewRequest(
		"POST",
		baseURI+pathShortener+"?key="+g.APIKey,
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return longURL, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Duration(g.Timeout) * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return longURL, err
	}
	defer res.Body.Close()
	type ResBody struct {
		ID string `json:"id"`
	}
	var resBody ResBody
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		fmt.Println(err)
	}
	return resBody.ID, nil
}

// Long ...
func (g *Googl) Long(shortURL string) (string, error) {
	params := map[string][]string{
		"shortUrl": {shortURL},
		"key":      {g.APIKey},
	}
	vals := url.Values(params)
	req, err := http.NewRequest("GET", baseURI, nil)
	if err != nil {
		return shortURL, err
	}
	req.URL.Path = path.Join(req.URL.Path, pathShortener)
	req.URL.RawQuery = vals.Encode()
	httpClient := &http.Client{Timeout: time.Duration(g.Timeout) * time.Second}
	res, err := httpClient.Do(req)
	if err != nil {
		return shortURL, err
	}
	defer res.Body.Close()
	type ResBody struct {
		LongURL string `json:"longUrl"`
	}
	var resBody ResBody
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return shortURL, err
	}
	return resBody.LongURL, nil
}
