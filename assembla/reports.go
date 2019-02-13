package assembla

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Assembla Report endpoint.
const customReportURL string = "https://api.assembla.com/v1/spaces/spaceID/tickets/custom_reports.json"

// ReportsService exposes all Report methods to the Client.
type ReportsService service

// Reports represents ther return object of Report methods.
type Reports struct {
	UserReports []ReportItem `json:"user_reports,omitempty"`
	TeamReports []ReportItem `json:"team_reports,omitempty"`
}

// ReportItem is a subset of the Report return object.
type ReportItem struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

// GetCustomReportsBySpaceID fetches all custom reports belonging to a given space.
// GET /v1/spaces/:space_id/tickets/custom_reports
// https://api-docs.assembla.cc/content/ref/tickets_custom_reports.html
func (s *ReportsService) GetCustomReportsBySpaceID(spaceID string) (reports Reports, err error) {
	url := strings.Replace(customReportURL, "spaceID", spaceID, -1)
	body, err := s.client.FetchRequestBody(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &reports)
	if err != nil {
		return reports, fmt.Errorf("Failed to unmarshal json (%s) --> %s", url, err)
	}

	return
}
