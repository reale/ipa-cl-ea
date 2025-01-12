package ipa

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type Error struct {
	Message string `json:"Error Message"`
}

type Client struct {
	APIKey   string
	BaseURL  string
	Endpoint string
}

func NewClient(apiKey string) *Client {
	c := &Client{}
	c.BaseURL = "https://www.indicepa.gov.it"
	c.Endpoint = "/public-ws/WS01_SFE_CF.php"
	c.APIKey = apiKey
	return c
}

func (c *Client) Query(params map[string]interface{}) (map[string]*json.RawMessage, error) {
	var om map[string]*json.RawMessage

	payload := url.Values{}

	payload.Add("AUTH_ID", c.APIKey)
	for k, v := range params {
		payload.Add(k, fmt.Sprint(v))
	}

	resp, err := http.PostForm(fmt.Sprintf("%s%s", c.BaseURL, c.Endpoint), payload)

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &om)
	if err != nil {
		return nil, err
	}
	if om["Error Message"] != nil {
		b, err := om["Error Message"].MarshalJSON()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(strings.Replace(string(b), "\"", "", -1))
	}
	return om, err
}
