package zendesk

import (
	"fmt"
	"io/ioutil"

	"github.com/go-resty/resty/v2"
	"gopkg.in/yaml.v2"
)

func FromToken(config ZendeskConfiguration) Client {
	username := fmt.Sprintf("%s/token", config.Email)

	restyClient := resty.New()
	restyClient.SetBasicAuth(username, config.Token)
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetHeader("Content-Type", "application/json")

	return Client{
		domain:     config.Domain,
		apiVersion: config.ApiVersion,
		client:     restyClient,
	}
}

func LoadConfiguration(path string) ZendeskConfiguration {
	zendeskConfiguration := ZendeskConfiguration{}
	zendeskConfigurationFile, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	err = yaml.Unmarshal(zendeskConfigurationFile, &zendeskConfiguration)

	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return zendeskConfiguration
}
