package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	ticketComments string = "https://api.assembla.com/v1/spaces/_space_id/tickets/_ticket_number/ticket_comments.json"
)

// Comment ...
//
// Assembla Docs: https://api-docs.assembla.cc/content/ref/ticket_comments_fields.html
type Comment struct {
	ID            int       `json:"id,omitempty"`
	Comment       string    `json:"comment,omitempty"`
	TicketID      int       `json:"ticket_id,omitempty"`
	UserID        string    `json:"user_id,omitempty"`
	CreatedOn     time.Time `json:"created_on,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	TicketChanges string    `json:"ticket_changes,omitempty"`
}

// GetTicketComments ...
// GET /v1/spaces/:space_id/tickets/:ticket_number/ticket_comments
// Assembla Docs: https://api-docs.assembla.cc/content/ref/ticket_comments_index.html
func (ac *AssemblaClient) GetTicketComments(spaceID string, ticketNumber int) (comments []Comment, err error) {
	r := strings.NewReplacer("_space_id", spaceID, "_ticket_number", strconv.Itoa(ticketNumber))
	url := r.Replace(ticketComments)
	page := 1

	for {
		params := fmt.Sprintf("?page=%x&per_page=25", page)
		body, err := ac.FetchRequestBody(url + params)
		if err != nil {
			if strings.Contains(err.Error(), "204") { // no more comments
				break
			}
			return nil, err
		}
		var newComments []Comment
		err = json.Unmarshal(body, &newComments)
		if err != nil {
			return nil, err
		}
		for _, comment := range newComments {
			comments = append(comments, comment)
		}
		page++
	}
	return
}
