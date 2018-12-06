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
	ID           string    `json:"id"`
	Login        string    `json:"login"`
	Name         string    `json:"name"`
	Picture      string    `json:"picture"`
	Email        string    `json:"email"`
	Organization string    `json:"organization"`
	Phone        string    `json:"phone"`
	IM           ProfileIM `json:"im"`
	IM2          ProfileIM `json:"im2"`
}

// ProfileIM is a subset of the User return object.
type ProfileIM struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// GetUserPicture returns the picture of the specified user.
// https://api.assembla.com/v1/users/_id_or_login/picture
func (ac *AssemblaClient) GetUserPicture() (pic []byte, err error) {

	url := strings.Replace(pictureURL, "_id_or_login", ac.user.ID, -1)
	pic, err = ac.FetchRequestBody(url)
	if err != nil {
		return
	}

	return
}

// GetUserPictureByID returns the picture of the given user.
// https://api.assembla.com/v1/users/_id_or_login/picture
func (ac *AssemblaClient) GetUserPictureByID(id string) (pic []byte, err error) {
	url := strings.Replace(pictureURL, "_id_or_login", id, -1)
	pic, err = ac.FetchRequestBody(url)
	if err != nil {
		return
	}

	return
}

// GetUserProfile returns the profile of the specified user.
// https://api.assembla.com/v1/users/_id_or_login.json
func (ac *AssemblaClient) GetUserProfile() (profile User, err error) {
	url := strings.Replace(profileURL, "_id_or_login", ac.user.ID, -1)
	body, err := ac.FetchRequestBody(url)

	err = json.Unmarshal(body, &profile)
	if err != nil {
		return
	}

	return
}

// GetUserProfileByID returns the profile of the given user.
// https://api.assembla.com/v1/users/_id_or_login.json
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
// https://api.assembla.com/v1/spaces/_space_id/users.json
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
