package assembla

import (
	"encoding/json"
	"net/http"
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
func (s *ReportsService) GetCustomReportsBySpaceID(spaceID string) (Reports, error) {
	url := strings.Replace(customReportURL, "spaceID", spaceID, -1)
	var reports Reports

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return reports, err
	}
	req.Header.Set("X-Api-Key", s.client.key)
	req.Header.Set("X-Api-Secret", s.client.secret)

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return reports, err
	}
	defer resp.Body.Close()

	// var body []byte
	// if resp.StatusCode == 200 {
	// 	body, err = ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// } else if resp.StatusCode == 204 {
	// 	return nil, errors.New("Failed Fetch! --> 204 No Content")
	// } else if resp.StatusCode == 404 {
	// 	return nil, errors.New("Failed Fetch! --> 404 Not Found")
	// }
	// return body, nil

	err = json.NewDecoder(resp.Body).Decode(reports)
	if err != nil {
		return reports, err
	}

	// body, err := s.client.FetchRequestBody(url)
	// if err != nil {
	// 	return reports, err
	// }

	// err = json.Unmarshal(body, &reports)
	// if err != nil {
	// 	return reports, err
	// }

	return reports, nil
}
