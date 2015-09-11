package config

import (
	"flag"
)

const (
	ParamConfigFile           = "configFile"
	ParamDatabaseUserId       = "database_user_id"
	ParamDatabaseUserPassword = "database_user_password"
	ParamDatabaseName         = "database_name"
)

type CmdParams struct {
	ConfigFile           string
	DatabaseUserId       string
	DatabaseUserPassword string
	DatabaseName         string
}

var defaultCmdParams = &CmdParams{
	ConfigFile:           ".bk_settings",
	DatabaseUserId:       "",
	DatabaseUserPassword: "",
	DatabaseName:         "",
}

func NewCmdParams() *CmdParams {

	// define the flags we are looking for
	configFilePtr := flag.String(ParamConfigFile, defaultCmdParams.ConfigFile, "location of config file for params")
	databaseUserIdPtr := flag.String(ParamDatabaseUserId, "", "connect to database as user")
	databaseUserPasswordPtr := flag.String(ParamDatabaseUserPassword, "", "password for database connection")
	databaseNamePtr := flag.String(ParamDatabaseName, "", "database to connect to")

	// parse cmdline
	flag.Parse()

	// create CmdParams and return
	cmdParams := new(CmdParams)
	cmdParams.ConfigFile = *configFilePtr
	cmdParams.DatabaseUserId = *databaseUserIdPtr
	cmdParams.DatabaseUserPassword = *databaseUserPasswordPtr
	cmdParams.DatabaseName = *databaseNamePtr

	return cmdParams

}
