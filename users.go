package main

import (
	"encoding/json"
	"strings"
)

const (
	userURL       string = "https://api.assembla.com/v1/user.json"
	profileURL    string = "https://api.assembla.com/v1/users/_id_or_login.json"
	pictureURL    string = "https://api.assembla.com/v1/users/_id_or_login/picture"
	spaceUsersURL string = "https://api.assembla.com/v1/spaces/_space_id/users.json"
)

// User ...
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

// ProfileIM ...
type ProfileIM struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// GetUser ... GET /v1/user
func (ac *AssemblaClient) GetUser() (newUser User, err error) {
	body, err := ac.FetchReqBody(userURL)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &newUser)
	if err != nil {
		return
	}

	return
}

// GetUserPicture ... GET /v1/users/[id_or_login]/picture
func (ac *AssemblaClient) GetUserPicture() (pic []byte, err error) {

	url := strings.Replace(pictureURL, "_id_or_login", ac.user.ID, -1)
	pic, err = ac.FetchReqBody(url)
	if err != nil {
		return
	}

	return
}

// GetUserPictureByID ... GET /v1/users/[id_or_login]/picture
func (ac *AssemblaClient) GetUserPictureByID(id string) (pic []byte, err error) {
	url := strings.Replace(pictureURL, "_id_or_login", id, -1)
	pic, err = ac.FetchReqBody(url)
	if err != nil {
		return
	}

	return
}

// GetUserProfile ... GET /v1/users/[id_or_login]
func (ac *AssemblaClient) GetUserProfile() (profile User, err error) {
	url := strings.Replace(profileURL, "_id_or_login", ac.user.ID, -1)
	body, err := ac.FetchReqBody(url)

	err = json.Unmarshal(body, &profile)
	if err != nil {
		return
	}

	return
}

// GetUserProfileByID ... GET /v1/users/[id_or_login]
func (ac *AssemblaClient) GetUserProfileByID(id string) (profile User, err error) {
	url := strings.Replace(profileURL, "_id_or_login", id, -1)
	body, err := ac.FetchReqBody(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		return
	}

	return
}

// GetUsersBySpaceID ... GET /v1/spaces/[space_id]/users
func (ac *AssemblaClient) GetUsersBySpaceID(id string) (users []User, err error) {
	url := strings.Replace(spaceUsersURL, "_space_id", id, -1)
	body, err := ac.FetchReqBody(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		return
	}

	return
}
