package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type auth struct {
	UserMail  string `json:"user_mail"`
	UserToken string `json:"user_token"`
	ProfileId string `json:"profile_id"`
}

type prefix struct {
	CardPrefix string `json:"card_prefix"`
	UrlPrefix  string `json:"url_prefix"`
}

type git struct {
	FeatureTag string `json:"feature_tag"`
	FixTag     string `json:"fix_tag"`
}

type config struct {
	Auth   auth   `json:"auth"`
	Prefix prefix `json:"prefixes"`
	Git    git    `json:"git"`
}

func GetConfig() (config, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		panic("Something is really wrong with your system, the HOME env don't exist")
	}

	file, err := ioutil.ReadFile(fmt.Sprintf("%s/.jira_config.json", home))

	if err != nil {
		return config{}, errors.Errorf("[GetConfig] -> Could not find the config file // %v", err)
	}

	var data config

	err = json.Unmarshal(file, &data)

	if err != nil {
		return config{}, errors.Errorf("[GetConfig] -> Could not parse the config file // %v", err)
	}

	return data, nil
}
