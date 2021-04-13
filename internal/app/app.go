package app

import (
	"bufio"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/midaef/emmet-client/configs"
	"github.com/midaef/emmet-client/internal/cli"
	"github.com/midaef/emmet-client/internal/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
	"strconv"
	"strings"
)

func Run(config *configs.Config) {
	opts := []grpc.DialOption {
		grpc.WithInsecure(),
	}
	host := fmt.Sprintf("%s:%s", config.Host, strconv.Itoa(config.Port))
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		grpclog.Fatalf("Fail to dial: %v", err)
	}

	defer conn.Close()

	tm.Clear()

	fmt.Println(tm.Color(tm.Bold("EMMET-CLIENT-CLI"), tm.GREEN) + "\n")

	clients := client.NewClients(conn)
	cli := cli.NewCLI(clients)

	if cli.UserInfo.AccessToken != "" {
		fmt.Println(tm.Bold("Connected to emmet-server successfully") + "\n")
	} else {
		fmt.Println(tm.Color(tm.Bold("Connected to emmet-server unsuccessfully"), tm.RED) + "\n")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		if cli.UserInfo.AccessToken != "" {
			terminal := fmt.Sprintf("%s@%s>>> ", cli.UserInfo.Login, host)
			fmt.Print(tm.Color(tm.Bold(terminal), tm.GREEN))

			command, _ := reader.ReadString('\n')
			command = strings.TrimSuffix(command, "\n")

			cli.CheckCommand(command)
		} else {
			break
		}
	}
}
