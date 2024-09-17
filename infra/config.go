package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Define the structure of your config
type Config struct {
	CICD struct {
		Env struct {
			Account string `yaml:"account"`
			Region  string `yaml:"region"`
		} `yaml:"env"`
		GithubTokenArn string `yaml:"githubTokenArn"`
		Repo           struct {
			Owner        string `yaml:"owner"`
			Repo         string `yaml:"repo"`
			Branch       string `yaml:"branch"`
			Path         string `yaml:"path"`
			PipelineName string `yaml:"pipelineName"`
			SentryToken  string `yaml:"sentryToken"`
			NpmToken     string `yaml:"npmToken"`
		} `yaml:"repo"`
	} `yaml:"cicd"`
}

func main() {
	// Fetch values from environment variables
	awsAccount := os.Getenv("AWS_ACCOUNT_ID")
	awsRegion := os.Getenv("AWS_REGION")

	if awsAccount == "" || awsRegion == "" {
		log.Fatalf("Missing required environment variables AWS_ACCOUNT_ID or AWS_REGION")
	}

	// Create the config struct
	var config Config
	config.CICD.Env.Account = awsAccount
	config.CICD.Env.Region = awsRegion
	config.CICD.GithubTokenArn = fmt.Sprintf("arn:aws:secretsmanager:%s:%s:secret:github/cicd-lax70A", awsRegion, awsAccount)
	config.CICD.Repo.Owner = "ckhine"
	config.CICD.Repo.Repo = "coaching"
	config.CICD.Repo.Branch = "main"
	config.CICD.Repo.Path = ""
	config.CICD.Repo.PipelineName = "main"

	// Marshal the config struct to YAML
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("Error marshaling config to YAML: %v", err)
	}

	// Save the YAML data to config.yml
	err = os.WriteFile("config.yml", yamlData, 0644)
	if err != nil {
		log.Fatalf("Error writing config.yml: %v", err)
	}

	fmt.Println("config.yml generated successfully")
}
