package cli

import "github.com/midaef/emmet-client/internal/client"

type Role struct {
	clients     *client.Clients
	accessToken string
}

func NewRoleCLI(clients *client.Clients, accessToken string) *Role {
	return &Role{
		clients: clients,
		accessToken: accessToken,
	}
}
