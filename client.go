package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// AssemblaClient ...
type AssemblaClient struct {
	key        string
	secret     string
	httpClient http.Client
	user       User
}

// FetchReqBody ...
func (ac *AssemblaClient) FetchReqBody(url string) (body []byte, err error) {
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

// Connect ...
func (ac *AssemblaClient) Connect(key string, secret string) (err error) {

	// setup
	ac.key = key
	ac.secret = secret

	// create http client
	ac.httpClient = http.Client{
		Timeout: time.Second * 2,
	}

	// test connection
	req, err := http.NewRequest(http.MethodGet, userURL, nil)
	if err != nil {
		return err
	}
	//req.Header.Set("User-Agent", "baileykg go-assembla client")
	req.Header.Set("X-Api-Key", key)
	req.Header.Set("X-Api-Secret", secret)

	// make request
	resp, err := ac.httpClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var profileUser User
	err = json.Unmarshal(body, &profileUser)
	if err != nil {
		return err
	}
	ac.user = profileUser

	return
}

// PrettyPrintJSON ...
func PrettyPrintJSON(body []byte) {
	var jsonOut bytes.Buffer
	err := json.Indent(&jsonOut, body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonOut.Bytes()))
}
