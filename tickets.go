package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	getSpaceTicketsURL string = "https://api.assembla.com/v1/spaces/_space_id/tickets.json"
)

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

// GetTicketsBySpaceAndReport retrieves all tickets belonging to a given space and report.
//
// Assembla Docs: https://api-docs.assembla.cc/content/ref/tickets_index.html
func (ac *AssemblaClient) GetTicketsBySpaceAndReport(reportID int, spaceID string) ([]Ticket, error) {
	url := strings.Replace(getSpaceTicketsURL, "_space_id", spaceID, -1)
	var allTickets []Ticket
	page := 1

	for {
		params := fmt.Sprintf("?report=%x&page=%x&per_page=100", reportID, page)
		body, err := ac.FetchRequestBody(url + params)
		if err != nil {
			if strings.Contains(err.Error(), "204") { // no more tickets
				break
			}
			return nil, err
		}

		var newTickets []Ticket
		err = json.Unmarshal(body, &newTickets)
		if err != nil {
			return nil, err
		}
		for _, ticket := range newTickets {
			allTickets = append(allTickets, ticket)
		}
		page++
	}

	return allTickets, nil
}
