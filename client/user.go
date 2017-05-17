package client

const (
	userUrlStr = "/user"
)

// UserService defines
type UserService struct {
	client *Client
}

// Sync sends request for user activation in k8s system
func (u *UserService) Sync(user User) error {
	req, err := u.client.NewRequest(postMethod, userUrlStr, user)
	if err != nil {
		return err
	}

	_, err = u.client.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

// User defines
type User struct {
	Name string `json:"name"`
}

// NewUser create a new User instance
func NewUser(username string) *User {
	return &User{
		Name: username,
	}
}
