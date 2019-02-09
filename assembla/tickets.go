package assembla

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	getSpaceTicketsURL    string = "https://api.assembla.com/v1/spaces/_space_id/tickets.json"
	getActiveTicketsURL   string = "https://api.assembla.com/v1/spaces/_space_id/tickets/my_active.json"
	getFollowedTicketsURL string = "https://api.assembla.com/v1/spaces/_space_id/tickets/my_followed.json"
)

// TicketsService ...
type TicketsService service

// Ticket ...
type Ticket struct {
	ID                 int       `json:"id,omitempty"`
	Number             int       `json:"number,omitempty"`
	Summary            string    `json:"summary,omitempty"`
	Description        string    `json:"description,omitempty"`
	Priority           int       `json:"priority,omitempty"`
	CompletedDate      time.Time `json:"completed_date,omitempty"`
	ComponentID        int       `json:"component_id,omitempty"`
	CreatedOn          time.Time `json:"created_on,omitempty"`
	PermissionType     int       `json:"permission_type,omitempty"`
	Importance         float64   `json:"importance,omitempty"`
	IsSTory            bool      `json:"is_story,omitempty"`
	MilestoneID        int       `json:"milestone_id,omitempty"`
	Tags               []string  `json:"tags,omitempty"`
	Followers          []string  `json:"followers,omitempty"`
	NotificationList   string    `json:"notification_list,omitempty"`
	SpaceID            string    `json:"space_id,omitempty"`
	State              int       `json:"state,omitempty"`
	Status             string    `json:"status,omitempty"`
	StoryImportance    int       `json:"story_importance,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	WorkingHours       float32   `json:"working_hours,omitempty"`
	Estimate           float32   `json:"estimate,omitempty"`
	TotalEstimate      float32   `json:"total_estimate,omitempty"`
	TotalInvestedHours float32   `json:"total_invested_hours,omitempty"`
	TotalWorkingHours  float32   `json:"total_working_hours,omitempty"`
	AssignedToID       string    `json:"assigned_to_id,omitempty"`
	ReporterID         string    `json:"reporter_id,omitempty"`
	HierarchyType      int       `json:"hierarchy_type,omitempty"`
	IsSupport          bool      `json:"is_support,omitempty"`
}

// GetActiveTicketsBySpace ...
// GET /v1/spaces/:space_id/tickets/my_active
// Assembla Docs: https://api-docs.assembla.cc/content/ref/tickets_my_active.html
func (s *TicketsService) GetActiveTicketsBySpace(spaceID string) ([]Ticket, error) {
	url := strings.Replace(getActiveTicketsURL, "_space_id", spaceID, -1)
	var activeTickets []Ticket
	page := 1

	for {
		params := fmt.Sprintf("?page=%x&per_page=25", page)
		body, err := s.client.FetchRequestBody(url + params)
		if err != nil {
			if strings.Contains(err.Error(), "204") { // no more tickets
				break
			}
			return nil, err
		}
		var newTickets []Ticket
		err = json.Unmarshal(body, &newTickets)
		if err != nil {
			return nil, errors.New("Failed to unmarshal json body")
		}
		for _, ticket := range newTickets {
			activeTickets = append(activeTickets, ticket)
		}
		page++
	}
	return activeTickets, nil
}

// GetFollowedTicketsBySpace ...
// GET /v1/spaces/[space_id]/tickets/my_followed
// Assembla Docs: https://api-docs.assembla.cc/content/ref/tickets_my_followed.html
func (s *TicketsService) GetFollowedTicketsBySpace(spaceID string) ([]Ticket, error) {
	url := strings.Replace(getFollowedTicketsURL, "_space_id", spaceID, -1)
	var followedTickets []Ticket
	page := 1

	for {
		params := fmt.Sprintf("?page=%x&per_page=25", page)
		body, err := s.client.FetchRequestBody(url + params)
		if err != nil {
			if strings.Contains(err.Error(), "204") {
				break
			}
			return nil, err
		}
		var fetchedTickets []Ticket
		err = json.Unmarshal(body, &fetchedTickets)
		if err != nil {
			return nil, errors.New("Failed to unmarshal json body")
		}
		for _, ticket := range fetchedTickets {
			followedTickets = append(followedTickets, ticket)
		}
		page++
	}
	return followedTickets, nil
}

// GetTicketsBySpaceAndReport retrieves all tickets belonging to a given space and report.
//
// Assembla Docs: https://api-docs.assembla.cc/content/ref/tickets_index.html
func (s *TicketsService) GetTicketsBySpaceAndReport(reportID int, spaceID string) ([]Ticket, error) {
	url := strings.Replace(getSpaceTicketsURL, "_space_id", spaceID, -1)
	var allTickets []Ticket
	page := 1

	for {
		params := fmt.Sprintf("?report=%x&page=%x&per_page=100", reportID, page)
		body, err := s.client.FetchRequestBody(url + params)
		if err != nil {
			if strings.Contains(err.Error(), "204") { // no more tickets
				break
			}
			return nil, err
		}

		var newTickets []Ticket
		err = json.Unmarshal(body, &newTickets)
		if err != nil {
			return nil, errors.New("Failed to unmarshal json body")
		}
		for _, ticket := range newTickets {
			allTickets = append(allTickets, ticket)
		}
		page++
	}

	return allTickets, nil
}
