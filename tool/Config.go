package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppPort  string         `json:"app_port"`
	AppHost  string         `json:"app_host"`
	Sms      SmsConfig      `json:"sms"`
	Database DatabaseConfig `json:"database"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}
func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("fmt 0192")
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err = decoder.Decode(&_cfg); err != nil {
		fmt.Println("fmt 0112")
		return nil, err
	}
	return _cfg, nil
}
