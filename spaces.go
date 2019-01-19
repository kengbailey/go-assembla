package main

import (
	"encoding/json"
	"strings"
	"time"
)

// Assembla Space endpoints.
const (
	spaceURL     string = "https://api.assembla.com/v1/spaces/_space_id/users.json"
	userSpaceURL string = "https://api.assembla.com/v1/spaces.json"
	spaceIDURL   string = "https://api.assembla.com/v1/spaces/_id.json"
)

// Space is the return object of all Assembla Space endpoints.
type Space struct {
	Status             int       `json:"status,omitempty"`
	BannerHeight       int       `json:"banner_height,omitempty"`
	Banner             string    `json:"banner,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	LastPayerChangedAt time.Time `json:"last_payer_changed_at,omitempty"`
	TeamTabRole        int       `json:"team_tab_role,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	Approved           bool      `json:"approved,omitempty"`
	TabsOrder          string    `json:"tabs_order,omitempty"`
	IsCommercial       bool      `json:"is_commercial,omitempty"`
	IsManager          bool      `json:"is_manager,omitempty"`
	TeamPermissions    int       `json:"team_permissions,omitempty"`
	CanJoin            bool      `json:"can_join,omitempty"`
	BannerText         string    `json:"banner_text,omitempty"`
	Restricted         bool      `json:"restricted,omitempty"`
	SharePermissions   bool      `json:"share_permissions,omitempty"`
	CanApply           bool      `json:"can_apply,omitempty"`
	IsVolunteer        bool      `json:"is_volunteer,omitempty"`
	PublicPermissions  int       `json:"public_permissions,omitempty"`
	WikiName           string    `json:"wiki_name,omitempty"`
	Name               string    `json:"name,omitempty"`
	Style              string    `json:"style,omitempty"`
	ParentID           string    `json:"parent_id,omitempty"`
	DefaultShowpage    string    `json:"default_showpage,omitempty"`
	Description        string    `json:"description,omitempty"`
	ID                 string    `json:"id,omitempty"`
	BannerLink         string    `json:"banner_link,omitempty"`
	CommercialFrom     time.Time `json:"commercial_from,omitempty"`
	RestrictedDate     string    `json:"restricted_date,omitempty"`
	WatcherPermissions int       `json:"watcher_permissions,omitempty"`
}

// GetUserSpaces returns a slice of spaces the specified user belongs to.
// https://api.assembla.com/v1/spaces.json
func (ac *AssemblaClient) GetUserSpaces() (spaces []Space, err error) {
	body, err := ac.FetchRequestBody(userSpaceURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &spaces)
	if err != nil {
		return
	}
	return
}

// GetSpaceByID returns a space given an id.
// https://api.assembla.com/v1/spaces/_id.json
func (ac *AssemblaClient) GetSpaceByID(id string) (space Space, err error) {
	url := strings.Replace(spaceIDURL, "_id", id, -1)
	body, err := ac.FetchRequestBody(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &space)
	if err != nil {
		return
	}
	return
}
