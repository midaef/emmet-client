package cli

import (
	"context"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/midaef/emmet-client/internal/api"
	"github.com/midaef/emmet-client/internal/client"
	"strconv"
	"strings"
)

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

func (c *Data) CreateValue() {
	fields := []string{"Key", "Value"}
	fieldsMap := Reader(fields)

	roles := []string{"How many roles can get the value?"}
	mapRoles := Reader(roles)
	i, err := strconv.Atoi(mapRoles[roles[0]])

	var fieldsRole []string

	for j := 0; j < i; j++ {
		fieldsRole = append(fieldsRole, strconv.Itoa(j + 1))
	}

	fieldsRoleMap := Reader(fieldsRole)

	var rolesList []string

	for _, f := range fieldsRole {
		rolesList = append(rolesList, fieldsRoleMap[f])
	}

	createValue := &api.CreateValueByAccessTokenRequest{
		AccessToken: c.accessToken,
		Key:         fieldsMap[fields[0]],
		Value:       fieldsMap[fields[1]],
		Roles:       strings.Join(rolesList, ","),
	}

	resp, err := c.clients.DataClient.CreateValueByAccessToken(context.Background(), createValue)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("%s", resp.Message)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}

func (c *Data) DeleteValue() {
	fields := []string{"Key"}
	fieldsMap := Reader(fields)

	deleteValue := &api.DeleteValueByAccessTokenRequest{
		AccessToken: c.accessToken,
		Key:         fieldsMap[fields[0]],
	}

	resp, err := c.clients.DataClient.DeleteValueByAccessToken(context.Background(), deleteValue)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("%s", resp.Message)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}

func (c *Data) GetValue() {
	fields := []string{"Key"}
	fieldsMap := Reader(fields)

	getValue := &api.GetValueByAccessTokenRequest{
		AccessToken: c.accessToken,
		Key:         fieldsMap[fields[0]],
	}

	resp, err := c.clients.DataClient.GetValueByAccessToken(context.Background(), getValue)
	if err != nil {
		error := fmt.Sprintf("%v\n", err)
		fmt.Println("\n" + tm.Color(tm.Bold(error), tm.RED) + "\n")

		return
	}

	message := fmt.Sprintf("Value: %s", resp.Value)
	fmt.Println("\n" + tm.Bold(message) + "\n")
}