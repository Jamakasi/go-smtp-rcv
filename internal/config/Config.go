package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	S_addr    string `yaml:"address"`
	S_type    string `yaml:"type"`
	S_domain  string `yaml:"domain"`
	S_version string `yaml:"version"`
}

type Config struct {
	Bind []Server
}

func NewConfig() *Config {
	var cfg Config

	return &cfg
}

func (c *Config) processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func (c *Config) Read(path string) {
	f, err := os.Open(path)
	if err != nil {
		c.processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		c.processError(err)
	}
}
func (c *Config) Write(path string) {
	yamlFile, err := yaml.Marshal(&c)
	if err != nil {
		c.processError(err)
	}
	f, err := os.Create(path)
	if err != nil {
		c.processError(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, string(yamlFile))
	if err != nil {
		c.processError(err)
	}
}

func (c *Config) GenerateExampleConfig() {
	s_tcp4_server := Server{S_addr: "127.0.0.1:2525",
		S_type:    "tcp4",
		S_domain:  "smtp.example.com",
		S_version: "rfc821"}
	s_tcp6_server := Server{S_addr: "[::]:2525",
		S_type:    "tcp6",
		S_domain:  "smtp.example.com",
		S_version: "rfc821"}

	c.Bind = []Server{s_tcp4_server, s_tcp6_server}
}
