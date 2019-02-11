package assembla

import (
	"encoding/json"
	"fmt"
	"strings"
)

const customReportURL string = "https://api.assembla.com/v1/spaces/spaceID/tickets/custom_reports.json"

// ReportsService ...
type ReportsService service

// Reports ...
type Reports struct {
	UserReports []ReportItem `json:"user_reports,omitempty"`
	TeamReports []ReportItem `json:"team_reports,omitempty"`
}

// ReportItem ...
type ReportItem struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

// GetCustomReportsBySpaceID ...
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
