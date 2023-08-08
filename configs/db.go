package configs

import (
	"fmt"
	"os"
	"strconv"
)

const (
	nameEnvKey     = "DB_NAME"
	hostEnvKey     = "DB_HOST"
	passwordEnvKey = "DB_PASS"
	portEnvKey     = "DB_PORT"
	usernameEnvKey = "DB_USER"
)

type DBConfig struct {
	Name     string
	Host     string
	Password string
	Port     int
	Username string
}

// GetDBConfig gets the configurations needed to connect to the db
// TODO: Add validation
// TODO: Use third party to get environment variables
func GetDBConfig() (*DBConfig, error) {
	var (
		dbCfg DBConfig
		err   error
	)

	dbCfg.Name = os.Getenv(nameEnvKey)
	dbCfg.Host = os.Getenv(hostEnvKey)
	dbCfg.Password = os.Getenv(passwordEnvKey)

	port := os.Getenv(portEnvKey)
	dbCfg.Port, err = strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("failed to convert db port to int: %v", err)
	}

	dbCfg.Username = os.Getenv(usernameEnvKey)

	return &dbCfg, nil
}
