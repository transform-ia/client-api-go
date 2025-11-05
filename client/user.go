package client

import (
	"fmt"

	"github.com/portainer/client-api-go/v2/pkg/client/users"
	"github.com/portainer/client-api-go/v2/pkg/models"
)

// ListUsers lists all users
func (c *PortainerClient) ListUsers() ([]*models.PortainereeUser, error) {
	params := users.NewUserListParams()
	resp, err := c.cli.Users.UserList(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return resp.Payload, nil
}

// CreateUser creates a new user.
func (c *PortainerClient) CreateUser(username, password string, role int64) (int64, error) {
	payload := &models.UsersUserCreatePayload{
		Username: &username,
		Password: &password,
	}
	payload.Role.PortainerUserRole = models.PortainerUserRole(role)

	params := users.NewUserCreateParams().WithBody(payload)
	resp, err := c.cli.Users.UserCreate(params, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return resp.Payload.ID, nil
}

// GetUser returns a user by ID
func (c *PortainerClient) GetUser(id int) (*models.PortainereeUser, error) {
	params := users.NewUserInspectParams().WithID(int64(id))
	resp, err := c.cli.Users.UserInspect(params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return resp.Payload, nil
}

// UpdateUserRole updates the role of a user.
func (c *PortainerClient) UpdateUserRole(id int, role int64) error {
	// Get current user details to fill required fields
	user, err := c.GetUser(id)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	payload := &models.UsersUserUpdatePayload{
		Username:    &user.Username,
		Password:    new(string), // Empty password when not changing it
		NewPassword: new(string), // Empty password when not changing it
		UseCache:    &user.UseCache,
	}
	payload.Role.PortainerUserRole = models.PortainerUserRole(role)

	params := users.NewUserUpdateParams().WithID(int64(id)).WithBody(payload)
	_, err = c.cli.Users.UserUpdate(params, nil)
	return err
}
