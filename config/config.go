package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type RuntimeConfig struct {
	DatabaseUserId       string `json:"database_user_id"`
	DatabaseUserPassword string `json:"database_user_password"`
	DatabaseName         string `json:"database_name"`
}

func (config *RuntimeConfig) override(params *CmdParams) {

	if params.DatabaseUserId != "" {
		config.DatabaseUserId = params.DatabaseUserId
	}
	if params.DatabaseUserPassword != "" {
		config.DatabaseUserPassword = params.DatabaseUserPassword
	}
	if params.DatabaseName != "" {
		config.DatabaseName = params.DatabaseName
	}
}

func setDefaultConfig(config *RuntimeConfig, params *CmdParams) {

	if config.DatabaseUserId == "" {
		config.DatabaseUserId = params.DatabaseUserId
	}
	if config.DatabaseUserPassword == "" {
		config.DatabaseUserPassword = params.DatabaseUserPassword
	}
	if config.DatabaseName == "" {
		config.DatabaseName = params.DatabaseName
	}
}

func Parse() *RuntimeConfig {

	// load the configuration values into config struct
	//   allow overriding of config entries on cmdline values
	cmdParams := NewCmdParams()

	// login to DB based on config object
	//   save results to singleton results store. store the resulting structs
	config, err := Load(cmdParams)
	if err != nil {
		panic(err)
	}

	return config
}

func Load(params *CmdParams) (*RuntimeConfig, error) {

	// load from config file if available
	// check that the config file exists, load if available
	f, err := os.Open(params.ConfigFile)
	if os.IsNotExist(err) {
		log.Println("Failed to find file:", params.ConfigFile)
		config := new(RuntimeConfig)
		setDefaultConfig(config, defaultCmdParams)
		config.override(params)
		if params.DatabaseName == "" || params.DatabaseUserId == "" || params.DatabaseUserPassword == "" {
			return nil, errors.New("Missing required configuration (Database, UserId, Password)")
		}
		return config, nil
	}
	defer f.Close()

	jsonByte, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	config := new(RuntimeConfig)
	if err := json.Unmarshal(jsonByte, &config); err != nil {
		return nil, err
	}
	setDefaultConfig(config, defaultCmdParams)
	config.override(params)
	if config.DatabaseName == "" || config.DatabaseUserId == "" || config.DatabaseUserPassword == "" {
		fmt.Println(params.DatabaseName, params.DatabaseUserId, params.DatabaseUserPassword)
		return nil, errors.New("Found config file, Missing required configuration (Database, UserId, Password)")
	}

	return config, nil
}

func BuildConnectionString(config *RuntimeConfig) string {

	return fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", config.DatabaseUserId, config.DatabaseUserPassword, config.DatabaseName)

}
