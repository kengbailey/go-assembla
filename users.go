package main

import (
	"encoding/json"
	"strings"
)

// Assembla User endpoints.
const (
	userURL       string = "https://api.assembla.com/v1/user.json"
	profileURL    string = "https://api.assembla.com/v1/users/_id_or_login.json"
	pictureURL    string = "https://api.assembla.com/v1/users/_id_or_login/picture"
	spaceUsersURL string = "https://api.assembla.com/v1/spaces/_space_id/users.json"
)

// User is the return object of all Assembla User endpoints.
type User struct {
	ID           string    `json:"id,omitempty"`
	Login        string    `json:"login,omitempty"`
	Name         string    `json:"name,omitempty"`
	Picture      string    `json:"picture,omitempty"`
	Email        string    `json:"email,omitempty"`
	Organization string    `json:"organization,omitempty"`
	Phone        string    `json:"phone,omitempty"`
	IM           ProfileIM `json:"im,omitempty"`
	IM2          ProfileIM `json:"im2,omitempty"`
}

// ProfileIM is a subset of the User return object.
type ProfileIM struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

// GetUser returns a minimal user struct with ID, name, login and phone.
// GET /v1/user
// Assembla Docs: https://api.assembla.com/v1/user.json
func (ac *AssemblaClient) GetUser() (user User, err error) {
	body, err := ac.FetchRequestBody(userURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return
	}
	return
}

// GetUserPicture returns the picture of the specified user.
// GET /v1/users/[id_or_login]/picture
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login/picture/
func (ac *AssemblaClient) GetUserPicture() ([]byte, error) {
	url := strings.Replace(pictureURL, "_id_or_login", ac.user.ID, -1)
	body, err := ac.FetchRequestBody(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetUserPictureByID returns the picture of the given user as a []byte.
// GET /v1/users/[id_or_login]/picture
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login/picture
func (ac *AssemblaClient) GetUserPictureByID(id string) ([]byte, error) {
	url := strings.Replace(pictureURL, "_id_or_login", id, -1)
	body, err := ac.FetchRequestBody(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetUserProfile returns the profile of the specified user.
// GET /v1/users/[id_or_login]
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login.json
func (ac *AssemblaClient) GetUserProfile() (profile User, err error) {
	url := strings.Replace(profileURL, "_id_or_login", ac.user.ID, -1)
	body, err := ac.FetchRequestBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return
	}
	return
}

// GetUserProfileByID returns the profile of the given user.
// GET /v1/users/[id_or_login]
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login.json
func (ac *AssemblaClient) GetUserProfileByID(id string) (profile User, err error) {
	url := strings.Replace(profileURL, "_id_or_login", id, -1)
	body, err := ac.FetchRequestBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return
	}
	return
}

// GetUsersBySpaceID returns a slice of users belonging to a given a space.
// GET /v1/spaces/[space_id]/users
// Assembla Docs: https://api.assembla.com/v1/spaces/_space_id/users.json
func (ac *AssemblaClient) GetUsersBySpaceID(id string) (users []User, err error) {
	url := strings.Replace(spaceUsersURL, "_space_id", id, -1)
	body, err := ac.FetchRequestBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return
	}
	return
}
