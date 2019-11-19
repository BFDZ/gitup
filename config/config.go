package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Repo struct {
	Dir    string `yaml:"dir"`
	Branch string `yaml:"branch"`
	Mode   string `yaml:"mode"`
	Time   string `yaml:"time"`
}

var Repos []Repo

func init() {
	conf := struct {
		Repo []Repo `yaml:"repo"`
	}{}
	if f, err := os.Open("./config.yml"); err != nil {
		panic(errors.New("can't find config.yml"))
	} else {
		err := yaml.NewDecoder(f).Decode(&conf)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("config loaded")

	Repos = conf.Repo
}
