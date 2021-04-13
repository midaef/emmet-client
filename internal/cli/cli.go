package cli

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
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
	DeleteUser()
}

type RoleCLI interface {
	CreateRole()
	DeleteRole()
}

type DataCLI interface {
	CreateValue()
	DeleteValue()
	GetValue()
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
		c.RoleCLI.CreateRole()
	case "create value":
		c.DataCLI.CreateValue()
	case "delete user":
		c.UserCLI.DeleteUser()
	case "delete role":
		c.RoleCLI.DeleteRole()
	case "delete value":
		c.DataCLI.DeleteValue()
	case "get value":
		c.DataCLI.GetValue()
	case "get token":
		token := fmt.Sprintf("Access token: %s", c.UserInfo.AccessToken)
		fmt.Println("\n" + tm.Bold(token) + "\n")
	case "help":
		c.help()
	default:
		fmt.Printf("\n" + tm.Color(tm.Bold("Command not found. Use "), tm.RED) +
			tm.Color(tm.Bold("help"), tm.GREEN) + "\n\n")
	}
}

func (c *CLI) help() {
	fmt.Printf("\n" + tm.Bold("Commands") + "\n")
	
	commands := []string{"create user", "create role", "create value", "delete user", "delete role", "delete value", "get value"}
	info := []string{"new user with role", "new role with permissions", "new value with role",
		"user will be deleted", "role will be deleted", "value will be deleted", "get value with key and permissions"}
	generateHelp(commands, info, len(commands))
}

func generateHelp(commands []string, info []string, fields int) {
	for i := 0; i < fields; i++ {
		fmt.Printf("\n" + tm.Color(tm.Bold(commands[i] + " - "), tm.GREEN) +
			tm.Bold(info[i]) + "\n")
	}

	fmt.Print("\n")
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




