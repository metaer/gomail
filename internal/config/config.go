package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const ConfigFilePath = "/etc/gomail/config.yaml"

type ServerConfig struct {
	SmtpServerDomain   string `yaml:"smtp_server_domain"`
	FromDomain         string `yaml:"from_domain"`
	Port               string `yaml:"port"`
	DkimPrivateKeyFile string `yaml:"dkim_private_key_file"`
	AddDkimSignature   bool   `yaml:"add_dkim_signature"`
	TlsCert            string `yaml:"tls_cert"`
	TlsKey             string `yaml:"tls_key"`
	DkimSelector       string `yaml:"dkim_selector"`
	Users              []struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"users"`
}

type DevClientConfig struct {
	DevRcptEmail string `yaml:"dev_rcpt_email"`
	FromDomain   string `yaml:"from_domain"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

func NewServerConfigFromYaml(configFile string) (*ServerConfig, error) {
	cfg := &ServerConfig{}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func NewDevClientConfigFromYaml(configFile string) (*DevClientConfig, error) {
	cfg := &DevClientConfig{}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
