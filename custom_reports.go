package main

import (
	"encoding/json"
	"strings"
)

const customReportURL string = "https://api.assembla.com/v1/spaces/spaceID/tickets/custom_reports.json"

// Reports ...
type Reports struct {
	UserReports []ReportItem `json:"user_reports"`
	TeamReports []ReportItem `json:"team_reports"`
}

// ReportItem ...
type ReportItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// GetCustomReportsBySpaceID ...
func (ac *AssemblaClient) GetCustomReportsBySpaceID(spaceID string) (reports Reports, err error) {
	url := strings.Replace(customReportURL, "spaceID", spaceID, -1)
	body, err := ac.FetchRequestBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &reports)
	if err != nil {
		return
	}
	return
}
