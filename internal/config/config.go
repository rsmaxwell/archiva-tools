package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/rsmaxwell/archiva/internal/archivaClient"
)

type Keep struct {
	MaximumNumber int `json:"maximumNumber"`
}

type Item struct {
	GroupId       string   `json:"groupId"`
	ArtifactId    string   `json:"artifactId"`
	Packaging     string   `json:"packaging"`
	RepositoryIds []string `json:"repositoryIds"`
	Keep          Keep     `json:"keep"`
}

type Config struct {
	Scheme       string                             `json:"scheme"`
	Host         string                             `json:"host"`
	Port         *int                               `json:"port"`
	Base         string                             `json:"base"`
	Users        []*archivaClient.User              `json:"users"`
	Cleanup      []*Item                            `json:"cleanup"`
	Repositories []*archivaClient.ManagedRepository `json:"repositories"`
}

const (
	DefaultConfigFile = "config.json"
)

// Open returns the configuration
func Open(configFileName string) (*Config, error) {

	bytearray, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	err = json.Unmarshal(bytearray, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) User(username string) *archivaClient.User {

	for _, user := range cfg.Users {
		if username == user.Username {
			return user
		}
	}

	return nil
}
