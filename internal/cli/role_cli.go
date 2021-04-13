package cli

import (
	"context"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/midaef/emmet-client/internal/api"
	"github.com/midaef/emmet-client/internal/client"
	"strconv"
)

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

func (c *Role) CreateRole() {
	fields := []string{"Role", "\nPermissions\n\nCreate user", "Create role", "Create value", "Delete user", "Delete role",
		"Delete value"}
	fieldsMap := Reader(fields)

	var boolFields []bool

	for _, f := range fields {
		ok, _ := strconv.ParseBool(fieldsMap[f])
		boolFields = append(boolFields, ok)
	}

	createRole := &api.CreateRoleByAccessTokenRequest{
		AccessToken: c.accessToken,
		Role:        fieldsMap[fields[0]],
		CreateUser:  boolFields[1],
		CreateRole:  boolFields[2],
		CreateValue: boolFields[3],
		DeleteUser:  boolFields[4],
		DeleteRole:  boolFields[5],
		DeleteValue: boolFields[6],
	}

	resp, err := c.clients.RoleClient.CreateRoleByAccessToken(context.Background(), createRole)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("%s", resp.Message)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}

func (c *Role) DeleteRole() {
	fields := []string{"Role"}
	fieldsMap := Reader(fields)

	deleteRole := &api.DeleteRoleByAccessTokenRequest{
		AccessToken: c.accessToken,
		Role:        fieldsMap[fields[0]],
	}

	resp, err := c.clients.RoleClient.DeleteRoleByAccessToken(context.Background(), deleteRole)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("%s", resp.Message)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}