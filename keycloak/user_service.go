package keycloak

import (
	"context"
	"net/url"
	"strings"
)

// UserService interacts with all user resources
type UserService service

// NewUserService returns a new user service for working with user resources
// in a realm.
func NewUserService(c *Client) *UserService {
	return &UserService{
		client: c,
	}
}

// Create creates a new user and returns the ID
func (us *UserService) Create(ctx context.Context, realm string, user *UserRepresentation) (string, error) {
	u := us.client.BaseURL
	u.Path += "/realms/{realm}/users"

	response, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetBody(user).
		SetResult(user).
		Post(u.String())

	if err = us.client.exec(response, err); err != nil {
		return "", err
	}

	location, err := url.Parse(response.Header().Get("Location"))

	if err != nil {
		return "", err
	}

	components := strings.Split(location.Path, "/")

	return components[len(components)-1], nil
}

// Get returns a user in a realm
func (us *UserService) Get(ctx context.Context, realm string, userID string) (*UserRepresentation, error) {

	u := us.client.BaseURL
	u.Path += "/realms/{realm}/users/{userID}"

	user := &UserRepresentation{}

	response, err := us.client.newRequest(ctx).
		SetPathParams(map[string]string{
			"realm":  realm,
			"userID": userID,
		}).
		SetResult(user).
		Get(u.String())

	if err = us.client.exec(response, err); err != nil {
		return nil, err
	}

	return user, nil
}

// Find returns users based on query params
// Params:
// - email
// - first
// - firstName
// - lastName
// - max
// - search
// - usename
func (us *UserService) Find(ctx context.Context, realm string, params map[string]string) ([]UserRepresentation, error) {

	u := us.client.BaseURL
	u.Path += "/realms/{realm}/users"

	user := []UserRepresentation{}

	response, err := us.client.newRequest(ctx).
		SetQueryParams(params).
		SetPathParams(map[string]string{
			"realm": realm,
		}).
		SetResult(&user).
		Get(u.String())

	if err = us.client.exec(response, err); err != nil {
		return nil, err
	}

	return user, nil
}