package cli

import "github.com/midaef/emmet-client/internal/client"

type Data struct {
	clients     *client.Clients
	accessToken string
}

func NewDataCLI(clients *client.Clients, accessToken string) *Data {
	return &Data{
		clients: clients,
		accessToken: accessToken,
	}
}