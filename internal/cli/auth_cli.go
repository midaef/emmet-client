package cli

import (
	"fmt"
	"context"
	tm "github.com/buger/goterm"
	"github.com/midaef/emmet-client/internal/api"
	"github.com/midaef/emmet-client/internal/client"
)

type Auth struct {
	clients     *client.Clients
	accessToken string
}

func NewAuthCLI(clients *client.Clients, accessToken string) *Auth {
	return &Auth{
		clients: clients,
		accessToken: accessToken,
	}
}

func (c *Auth) AuthWithCredentials() string {
	fields := []string{"Login", "Password"}
	fieldsMap := Reader(fields)

	authRequest := &api.AuthWithCredentialsRequest{
		Login: fieldsMap[fields[0]],
		Password: fieldsMap[fields[1]],
	}

	resp, err := c.clients.AuthClient.AuthWithCredentials(context.Background(), authRequest)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return ""
	}

	token := fmt.Sprintf("Access token: %s", resp.AccessToken)
	fmt.Println("\n" + tm.Bold(token) + "\n")

	c.accessToken = resp.AccessToken

	return authRequest.Login
}