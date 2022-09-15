package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	//var body []byte
	var resp *http.Response

	if c.Auth.apikey == "" {
		return nil, fmt.Errorf("apikey missing from tanzu provider config")
	}
	// rb, err := json.Marshal(c.Auth)
	// if err != nil {
	// 	return nil, err
	// }

	data := url.Values{}
	data.Set("refresh_token", "vsoq2sBIO1QFA-d9pDgQK7RuVk-Rxie05UZsiN0uVGwR2VCCGvzLj5XP6d0UH8nm")

	req, err := http.NewRequest("POST", c.AuthURL, strings.NewReader(data.Encode()))
	if err == nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		debug(httputil.DumpRequestOut(req, true))
		resp, err = (&http.Client{}).Do(req)
	}

	if err == nil {
		defer resp.Body.Close()
		debug(httputil.DumpResponse(resp, true))
		body, err := ioutil.ReadAll(resp.Body)
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}
		return &ar, nil
	}

	return nil, nil
}

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}
