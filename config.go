package main

import (
	"io/ioutil"
	"log"
	"os"

	"path/filepath"

	"gopkg.in/yaml.v2"
)

const credsFileName string = "credentials.yml"

type credentials struct {
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
	return filepath.Join(curDir, credsFileName)
}

func (c *credentials) loadCredentials() *credentials {
	credentialsFileName := getCredsFilePath()

	yamlFile, err := ioutil.ReadFile(credentialsFileName)
	if err != nil {
		ErrorLogger.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		ErrorLogger.Fatalf("Unmarshal: %v", err)
	}

	return c
}
