package main

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	spaceURL     string = "https://api.assembla.com/v1/spaces/_space_id/users.json"
	userSpaceURL string = "https://api.assembla.com/v1/spaces.json"
	spaceIDURL   string = "https://api.assembla.com/v1/spaces/_id.json"
)

// Space ...
type Space struct {
	Status             int       `json:"status"`
	BannerHeight       int       `json:"banner_height"`
	Banner             string    `json:"banner"`
	UpdatedAt          time.Time `json:"updated_at"`
	LastPayerChangedAt time.Time `json:"last_payer_changed_at"`
	TeamTabRole        int       `json:"team_tab_role"`
	CreatedAt          time.Time `json:"created_at"`
	Approved           bool      `json:"approved"`
	TabsOrder          string    `json:"tabs_order"`
	IsCommercial       bool      `json:"is_commercial"`
	IsManager          bool      `json:"is_manager"`
	TeamPermissions    int       `json:"team_permissions"`
	CanJoin            bool      `json:"can_join"`
	BannerText         string    `json:"banner_text"`
	Restricted         bool      `json:"restricted"`
	SharePermissions   bool      `json:"share_permissions"`
	CanApply           bool      `json:"can_apply"`
	IsVolunteer        bool      `json:"is_volunteer"`
	PublicPermissions  int       `json:"public_permissions"`
	WikiName           string    `json:"wiki_name"`
	Name               string    `json:"name"`
	Style              string    `json:"style"`
	ParentID           string    `json:"parent_id"`
	DefaultShowpage    string    `json:"default_showpage"`
	Description        string    `json:"description"`
	ID                 string    `json:"id"`
	BannerLink         string    `json:"banner_link"`
	CommercialFrom     time.Time `json:"commercial_from"`
	RestrictedDate     string    `json:"restricted_date"`
	WatcherPermissions int       `json:"watcher_permissions"`
}

// GetUserSpaces ... GET /v1/spaces
func (ac *AssemblaClient) GetUserSpaces() (spaces []Space, err error) {
	body, err := ac.FetchReqBody(userSpaceURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &spaces)
	if err != nil {
		return
	}
	return
}

// GetSpaceByID ... GET /v1/spaces/:id_or_wiki_name
func (ac *AssemblaClient) GetSpaceByID(id string) (space Space, err error) {
	url := strings.Replace(spaceIDURL, "_id", id, -1)
	body, err := ac.FetchReqBody(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &space)
	if err != nil {
		return
	}
	return
}

//
