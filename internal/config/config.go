package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Settings struct {
	Database    DatabaseSettings    `yaml:"database"`
	Application ApplicationSettings `yaml:"application"`
}

type ProductionSettings struct {
	Database    ProductionDatabaseSettings    `yaml:"database"`
	Application ProductionApplicationSettings `yaml:"application"`
}

type ApplicationSettings struct {
	Port    uint16 `yaml:"port"`
	Host    string `yaml:"host"`
	BaseUrl string `yaml:"baseurl"`
}

type ProductionApplicationSettings struct {
	Port uint16 `yaml:"port"`
	Host string `yaml:"host"`
}

type DatabaseSettings struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Port         uint16 `yaml:"port"`
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"databasename"`
	RequireSSL   bool   `yaml:"requiressl"`
}

type ProductionDatabaseSettings struct {
	RequireSSL bool `yaml:"requiressl"`
}

func (s *DatabaseSettings) WithoutDb() string {
	var sslMode string
	if s.RequireSSL {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s",
		s.Host, s.Port, s.Username, s.Password, sslMode)
}

func (s *DatabaseSettings) WithDb() string {
	return fmt.Sprintf("%s dbname=%s", s.WithoutDb(), s.DatabaseName)
}

type Environment string

const (
	EnvironmentLocal      Environment = "local"
	EnvironmentProduction Environment = "production"
)

func GetSettings(path string) (*Settings, error) {
	var environment Environment
	if env, ok := os.LookupEnv("APP_ENVIRONMENT"); ok {
		environment = Environment(env)
	} else {
		environment = "local"
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	if environment == EnvironmentLocal {
		var settings Settings

		v.SetConfigName(string(environment))

		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read %s configuration: %w", string(environment), err)
		}

		if err := v.Unmarshal(&settings); err != nil {
			return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
		}

		return &settings, nil
	} else {
		var prodSettings ProductionSettings

		v.SetConfigName(string(environment))

		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read %s configuration: %w", string(environment), err)
		}

		if err := v.Unmarshal(&prodSettings); err != nil {
			return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
		}

		appPrefix := "APP_"
		dbPrefix := fmt.Sprintf("%sDATABASE__", appPrefix)
		applicationPrefix := fmt.Sprintf("%sAPPLICATION__", appPrefix)

		dbPortStr := os.Getenv(fmt.Sprintf("%sPORT", dbPrefix))
		dbPortInt, err := strconv.Atoi(dbPortStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse db port: %w", err)
		}

		return &Settings{
			Database: DatabaseSettings{
				Username:     os.Getenv(fmt.Sprintf("%sUSERNAME", dbPrefix)),
				Password:     os.Getenv(fmt.Sprintf("%sPASSWORD", dbPrefix)),
				Host:         os.Getenv(fmt.Sprintf("%sHOST", dbPrefix)),
				Port:         uint16(dbPortInt),
				DatabaseName: os.Getenv(fmt.Sprintf("%sDATABASE_NAME", dbPrefix)),
				RequireSSL:   prodSettings.Database.RequireSSL,
			},
			Application: ApplicationSettings{
				Port:    prodSettings.Application.Port,
				Host:    prodSettings.Application.Host,
				BaseUrl: os.Getenv(fmt.Sprintf("%sBASE_URL", applicationPrefix)),
			},
		}, nil
	}
}
