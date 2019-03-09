package assembla

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	getAllMilestonesURL string = "https://api.assembla.com/v1/spaces/_space_id/milestones/all.json"
)

// MilestonesService exposes all Milestone methods to the client.
type MilestonesService service

// Milestone ...
// https://api-docs.assembla.cc/content/ref/milestones_fields.html
type Milestone struct {
	ID int `json:"id,omitempty"`
	//StartDate time.Time `json:"start_date,omitempty"`
	//DueDate   time.Time `json:"due_date,omitempty"`
	Budget float32 `json:"budget,omitempty"`
	Title  string  `json:"title,omitempty"`
	UserID string  `json:"user_id,omitempty"`
	//CreatedAt          time.Time `json:"created_at,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	SpaceID     string `json:"space_id,omitempty"`
	Description string `json:"description,omitempty"`
	IsCompleted bool   `json:"is_completed,omitempty"`
	//CompletedDate      time.Time `json:"completed_date,omitempty"`
	//UpdatedAt          time.Time `json:"updated_at,omitempty"`
	UpdatedBy          string `json:"updated_by,omitempty"`
	ReleaseLevel       int    `json:"release_level,omitempty"`
	ReleaseNotes       string `json:"release_notes,omitempty"`
	PlannerType        int    `json:"planner_type,omitempty"`
	PrettyReleaseLevel string `json:"pretty_release_level,omitempty"`
}

// GetAllMilestones ...
//
// Assembla Docs: https://api-docs.assembla.cc/content/ref/milestones_all.html
func (s *MilestonesService) GetAllMilestones(spaceID string) ([]Milestone, error) {
	url := strings.Replace(getAllMilestonesURL, "_space_id", spaceID, -1)
	var allMilestones []Milestone
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
		var newMilestones []Milestone
		err = json.Unmarshal(body, &newMilestones)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal milestone json (%s) --> %s", url+params, err)
		}
		for _, milestone := range newMilestones {
			allMilestones = append(allMilestones, milestone)
		}
		page++
	}
	return allMilestones, nil
}
