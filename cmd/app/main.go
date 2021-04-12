package main

import (
	"flag"
	"github.com/midaef/emmet-client/configs"
	"github.com/midaef/emmet-client/internal/app"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

const defaultConfigPath = "./configs/"

const defaultConfigName = "default_config.yaml"

const defaultServerPort = 65000

const defaultServerHost = "localhost"

var configName string

var configPath string

var (
	host string
	port int
)

func init() {
	flag.StringVar(&configName, "config-name", defaultConfigName, "config name")
	flag.StringVar(&configPath, "config-path", defaultConfigPath, "config path")
	flag.StringVar(&host, "server-host", defaultServerHost, "server port")
	flag.IntVar(&port, "server-port", defaultServerPort, "server port")
}

func main()  {
	flag.Parse()

	config, err := getConfig(configName)
	if config == nil || err != nil{
		config = &configs.Config{
			Host: host,
			Port: port,
		}
	}

	app.Run(config)
}

func getConfig(name string) (*configs.Config, error) {
	configPath = configPath + name

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config *configs.Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}