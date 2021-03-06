package assembla

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Assembla User endpoints.
const (
	userURL       string = "https://api.assembla.com/v1/user.json"
	profileURL    string = "https://api.assembla.com/v1/users/_id_or_login.json"
	pictureURL    string = "https://api.assembla.com/v1/users/_id_or_login/picture"
	spaceUsersURL string = "https://api.assembla.com/v1/spaces/_space_id/users.json"
)

// UsersService exposes all User methods to the Client.
type UsersService service

// User represents the return object of most Assembla User methods.
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
func (s *UsersService) GetUser() (user User, err error) {
	body, err := s.client.FetchRequestBody(userURL)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, fmt.Errorf("Failed to unmarshal json (%s) --> %s", userURL, err)
	}
	return
}

// GetUserPicture fetches the []byte picture of the authenticated user.
// GET /v1/users/[id_or_login]/picture
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login/picture/
func (s *UsersService) GetUserPicture() ([]byte, error) {
	url := strings.Replace(pictureURL, "_id_or_login", s.client.user.ID, -1)
	body, err := s.client.FetchRequestBody(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetUserPictureByID returns the picture of the given user as a []byte.
// GET /v1/users/[id_or_login]/picture
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login/picture
func (s *UsersService) GetUserPictureByID(id string) ([]byte, error) {
	url := strings.Replace(pictureURL, "_id_or_login", id, -1)
	body, err := s.client.FetchRequestBody(url)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetUserProfile fetches the profile of the authenticated user.
// GET /v1/users/[id_or_login]
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login.json
func (s *UsersService) GetUserProfile() (profile User, err error) {
	url := strings.Replace(profileURL, "_id_or_login", s.client.user.ID, -1)
	body, err := s.client.FetchRequestBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return profile, fmt.Errorf("Failed to unmarshal json (%s) --> %s", url, err)
	}
	return
}

// GetUserProfileByID fetches the profile of the given user.
// GET /v1/users/[id_or_login]
// Assembla Docs: https://api.assembla.com/v1/users/_id_or_login.json
func (s *UsersService) GetUserProfileByID(id string) (profile User, err error) {
	url := strings.Replace(profileURL, "_id_or_login", id, -1)
	body, err := s.client.FetchRequestBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return profile, fmt.Errorf("Failed to unmarshal json (%s) --> %s", url, err)
	}
	return
}

// GetUsersBySpaceID fetches a slice of Users belonging to a given a space.
// GET /v1/spaces/[space_id]/users
// Assembla Docs: https://api.assembla.com/v1/spaces/_space_id/users.json
func (s *UsersService) GetUsersBySpaceID(id string) (users []User, err error) {
	url := strings.Replace(spaceUsersURL, "_space_id", id, -1)
	body, err := s.client.FetchRequestBody(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return users, fmt.Errorf("Failed to unmarshal json (%s) --> %s", url, err)
	}
	return
}
