package assembla

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client is the client to Assembla's REST API v1.
// You can start a session by calling NewClient()
// This connects to assembla...
type Client struct {
	key        string
	secret     string
	httpClient http.Client
	user       User

	common service // reusable common service for initializing all client services

	// Client services
	Reports    *ReportsService
	Spaces     *SpacesService
	Tickets    *TicketsService
	Users      *UsersService
	Comments   *CommentsService
	Milestones *MilestonesService
}

type service struct {
	client *Client
}

// NewClient creates a new Client, given a key and secret for authentication.
// Basic user details are retrieved using the connect method.
func NewClient(key string, secret string) *Client {
	client := &Client{key: key, secret: secret}
	client.httpClient = http.Client{
		Timeout: time.Second * 2,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: false,
		},
	}

	client.common.client = client
	client.Reports = (*ReportsService)(&client.common)
	client.Spaces = (*SpacesService)(&client.common)
	client.Tickets = (*TicketsService)(&client.common)
	client.Users = (*UsersService)(&client.common)
	client.Comments = (*CommentsService)(&client.common)
	client.Milestones = (*MilestonesService)(&client.common)

	return client
}

// FetchRequestBody is used by endpoint methods to make a request, given a url.
// it returns the response body as []byte if request is successful.
func (ac *Client) FetchRequestBody(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", ac.key)
	req.Header.Set("X-Api-Secret", ac.secret)

	resp, err := ac.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body []byte
	if resp.StatusCode == 200 {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
	} else if resp.StatusCode == http.StatusNoContent {
		return nil, errors.New("Failed Fetch! --> 204 No Content")
	} else if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("Failed Fetch! --> 404 Not Found (%s)", url)
	}
	return body, nil
}

// connect fetches user details for a newly created Client.
func (ac *Client) connect(key string, secret string) (err error) {

	body, err := ac.FetchRequestBody(userURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &ac.user)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal json (%s) --> %s", userURL, err)
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
