package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"path/filepath"

	"github.com/loganballard/odds_briefing/logger"
	"gopkg.in/yaml.v2"
)

var credsFileName string = filepath.Join("credentials", "credentials.yml")

type Credentials struct {
	OddsApiKey       string `yaml:"odds_api_key"`
	TwilioSid        string `yaml:"twilio_sid"`
	TwilioAuthKey    string `yaml:"twilio_auth_key"`
	TwilioNumberFrom string `yaml:"twilio_number_from"`
	TwilioNumberTo   string `yaml:"twilio_number_to"`
}

func getCredsFilePath() string {
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if os.Getenv("CI") == "true" {
		return filepath.Join(curDir, credsFileName+".example")
	}
	return filepath.Join(curDir, credsFileName)
}

func (c *Credentials) loadCredentialsCI() *Credentials {
	c.OddsApiKey = os.Getenv("CI_ODDS_KEY")
	c.TwilioAuthKey = os.Getenv("CI_TWILIO_AUTH_KEY")
	c.TwilioSid = os.Getenv("CI_TWILIO_SID")
	c.TwilioNumberTo = os.Getenv("CI_TWILIO_NUMBER_TO")
	c.TwilioNumberFrom = os.Getenv("CI_TWILIO_NUMBER_FROM")
	return c
}

func (c *Credentials) LoadCredentials() *Credentials {
	credentialsFileName := getCredsFilePath()

	if os.Getenv("CI") == "true" {
		c.loadCredentialsCI()
		return c
	}

	yamlFile, err := ioutil.ReadFile(credentialsFileName)
	if err != nil {
		logger.Error(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		logger.Error(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}

	return c
}
