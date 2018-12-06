package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// AssemblaClient is the client to Assembla's REST API v1.
// You can start a session by calling NewClient()
// This connects to assembla...
type AssemblaClient struct {
	key        string
	secret     string
	httpClient http.Client
	user       User
}

// NewClient creates a new AssemblaClient, given a key and secret for authentication.
// Basic user details are retrieved using the connect method.
func NewClient(key string, secret string) (*AssemblaClient, error) {
	client := AssemblaClient{key: key, secret: secret}
	client.httpClient = http.Client{Timeout: time.Second * 2}
	err := client.connect(key, secret)
	if err != nil {
		return nil, err
	}
	return &client, err
}

// FetchRequestBody is used by endpoint methods to make a request, given a url.
// it returns the response body as []byte if request is successful.
func (ac *AssemblaClient) FetchRequestBody(url string) (body []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("X-Api-Key", ac.key)
	req.Header.Set("X-Api-Secret", ac.secret)

	resp, err := ac.httpClient.Do(req)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// connect fetches user details for a newly created AssemblaClient.
func (ac *AssemblaClient) connect(key string, secret string) (err error) {

	body, err := ac.FetchRequestBody(userURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ac.user)
	if err != nil {
		return
	}
	return
}

// PrettyPrintJSON prints json with indents for readability.
// Takes a byte slice, usually a request body.
func PrettyPrintJSON(body []byte) {
	var jsonOut bytes.Buffer
	_ = json.Indent(&jsonOut, body, "", "  ")
	fmt.Println(string(jsonOut.Bytes()))
}
