package client

import (
	"github.com/midaef/emmet-client/internal/api"
	"google.golang.org/grpc"
)

type Clients struct {
	AuthClient api.AuthClient
	DataClient api.DataClient
	RoleClient api.RoleClient
	UserClient api.UserClient
}

func NewClients(conn grpc.ClientConnInterface) *Clients {
	authClient := api.NewAuthClient(conn)
	dataClient := api.NewDataClient(conn)
	roleClient := api.NewRoleClient(conn)
	userClient := api.NewUserClient(conn)

	return &Clients{
		AuthClient: authClient,
		DataClient: dataClient,
		RoleClient: roleClient,
		UserClient: userClient,
	}
}


