package server

import (
	"encoding/json"
	"os"
)

type config struct {
	Addr       string `json:"addr"`
	MaxClients int    `json:"maxClients"`
}

type Configurator struct {
	configPath    string
	configuration *config
}

func NewConfigurator(configPath string) *Configurator {
	return &Configurator{
		configPath: configPath,
	}
}

func (c *Configurator) Load() error {
	file, err := os.Open(c.configPath)
	if err != nil {
		return err
	}
	fileStat, err := file.Stat()
	if err != nil {
		return err
	}

	configSize := fileStat.Size()

	configData := make([]byte, configSize)
	tempConfig := &config{}

	file.Read(configData)

	err = json.Unmarshal(configData, tempConfig)
	if err != nil {
		return err
	}

	c.configuration = tempConfig
	return nil
}

func (c *Configurator) Get() *config {
	return c.configuration
}

func (c *Configurator) Save() error {
	return nil
}
