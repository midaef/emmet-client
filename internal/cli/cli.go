package cli

import (
	"bufio"
	"fmt"
	"github.com/midaef/emmet-client/internal/client"
	"os"
	"strings"
)

type AuthCLI interface {
	AuthWithCredentials() string
}

type UserCLI interface {
	UpdatePassword()
	CreateUser()
}

type RoleCLI interface {

}

type DataCLI interface {

}

type CLI struct {
	AuthCLI  AuthCLI
	UserCLI  UserCLI
	RoleCLI  RoleCLI
	DataCLI  DataCLI

	UserInfo *UserInfo
}

type UserInfo struct {
	AccessToken string
	Login       string
}

func NewCLI(clients *client.Clients) *CLI {
	authCLI := NewAuthCLI(clients, "")
	login := authCLI.AuthWithCredentials()

	userCLI := NewUserCLI(clients, authCLI.accessToken)
	roleCLI := NewRoleCLI(clients, authCLI.accessToken)
	dataCLI := NewDataCLI(clients, authCLI.accessToken)

	return &CLI{
		AuthCLI: authCLI,
		UserCLI: userCLI,
		RoleCLI: roleCLI,
		DataCLI: dataCLI,

		UserInfo: &UserInfo{
			AccessToken: authCLI.accessToken,
			Login:       login,
		},
	}
}

func (c *CLI) CheckCommand(command string) {
	switch command {
	case "exit":
		c.UserInfo.AccessToken = ""
	case "logout":
		c.UserInfo.AccessToken = ""
	case "update password":
		c.UserCLI.UpdatePassword()
	case "create user":
		c.UserCLI.CreateUser()
	case "create role":
	case "create value":
	case "delete user":
	case "delete role":
	case "delete value":
	case "get value":
	case "get token":
		fmt.Printf("Access token: %s\n", c.UserInfo.AccessToken)
	case "help":
	default:
		fmt.Printf("Command not found. Use help\n")
	}
}

func Reader(fields []string) map[string]string {
	reader := bufio.NewReader(os.Stdin)
	fieldsMap := make(map[string]string)

	for _, f := range fields {
		fmt.Print(f + ": ")

		field, _ := reader.ReadString('\n')
		field = strings.TrimSuffix(field, "\n")

		fieldsMap[f] = field
	}

	return fieldsMap
}




