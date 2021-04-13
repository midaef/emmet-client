package cli

import (
	"context"
	"fmt"
	tm "github.com/buger/goterm"

	"github.com/midaef/emmet-client/internal/api"
	"github.com/midaef/emmet-client/internal/client"
)

type User struct {
	clients     *client.Clients
	accessToken string
}

func NewUserCLI(clients *client.Clients, accessToken string) *User {
	return &User{
		clients: clients,
		accessToken: accessToken,
	}
}

func (c *User) UpdatePassword() {
	field := []string{"New password"}
	fieldsMap := Reader(field)

	updatePassword := &api.UpdatePasswordByAccessTokenRequest{
		AccessToken: c.accessToken,
		Password: fieldsMap[field[0]],
	}

	resp, err := c.clients.UserClient.UpdatePasswordByAccessToken(context.Background(), updatePassword)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("%s", resp.Message)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}

func (c *User) CreateUser() {
	fields := []string{"Login", "Password", "Role"}
	fieldsMap := Reader(fields)

	createUser := &api.CreateUserByAccessTokenRequest{
		AccessToken: c.accessToken,
		Login: fieldsMap[fields[0]],
		Password: fieldsMap[fields[1]],
		Role: fieldsMap[fields[2]],
	}

	resp, err := c.clients.UserClient.CreateUserByAccessToken(context.Background(), createUser)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("%s", resp.Message)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}

func (c *User) DeleteUser() {
	fields := []string{"User login"}
	fieldsMap := Reader(fields)

	deleteUser := &api.DeleteUserByAccessTokenRequest{
		AccessToken: c.accessToken,
		Login: fieldsMap[fields[0]],
	}

	resp, err := c.clients.UserClient.DeleteUserByAccessToken(context.Background(), deleteUser)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("%s", resp.Message)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}