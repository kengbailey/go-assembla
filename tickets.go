package main

import (
	"time"
)

const (
	getSpaceTicketsURL string = "https://api.assembla.com/v1/spaces/_space_id/tickets.json"
)

// Ticket ...
type Ticket struct {
	ID                 int       `json:"id"`
	Number             int       `json:"number"`
	Summary            string    `json:"summary"`
	Description        string    `json:"description"`
	Priority           int       `json:"priority"`
	CompletedDate      time.Time `json:"completed_date"`
	ComponentID        int       `json:"component_id"`
	CreatedOn          time.Time `json:"created_on"`
	PermissionType     int       `json:"permission_type"`
	Importance         float64   `json:"importance"`
	IsSTory            bool      `json:"is_story"`
	MilestoneID        int       `json:"milestone_id"`
	Tags               []string  `json:"tags"`
	Followers          []string  `json:"followers"`
	NotificationList   string    `json:"notification_list"`
	SpaceID            string    `json:"space_id"`
	State              int       `json:"state"`
	Status             string    `json:"status"`
	StoryImportance    int       `json:"story_importance"`
	UpdatedAt          time.Time `json:"updated_at"`
	WorkingHours       float32   `json:"working_hours"`
	Estimate           float32   `json:"estimate"`
	TotalEstimate      float32   `json:"total_estimate"`
	TotalInvestedHours float32   `json:"total_invested_hours"`
	TotalWorkingHours  float32   `json:"total_working_hours"`
	AssignedToID       string    `json:"assigned_to_id"`
	ReporterID         string    `json:"reporter_id"`
	HierarchyType      int       `json:"hierarchy_type"`
	IsSupport          bool      `json:"is_support"`
	//CustomFields       string    `json:"custom_fields"`
}

// GetTicketsBySpaceAndReport ...
func (ac *AssemblaClient) GetTicketsBySpaceAndReport(reportID int, spaceID string) (tickets []Ticket, err error) {
	// retrieve all tickets belonging to user and report
	return
}

// // GetUserTicketsBySpaceID ...
// func (ac *AssemblaClient) GetUserTicketsBySpaceID(id string, num int) (Tickets []Ticket, err error) {
// 	// craft new url using params to fetch

// 	url := strings.Replace(getSpaceTicketsURL, "_space_id", id, -1)
// 	paramURL := url + "?report=4264781&page=1"

// 	// fetch tickets
// 	body, err := ac.FetchRequestBody(paramURL)
// 	if err != nil {
// 		return
// 	}

// 	err = json.Unmarshal(body, &Tickets)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(Tickets[0].Summary)

// 	return
// }
