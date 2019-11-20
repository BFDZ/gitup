package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Repo struct {
	Dir    string `yaml:"dir"`
	Branch string `yaml:"branch"`
	Mode   string `yaml:"mode"`
	Time   string `yaml:"time"`
}

var Repos []Repo
var ListenAddr string

func init() {
	conf := struct {
		Repo       []Repo `yaml:"repo"`
		ListenAddr string `yaml:"listen"`
	}{}
	if f, err := os.Open("./config.yml"); err != nil {
		panic(errors.New("can't find config.yml"))
	} else {
		err := yaml.NewDecoder(f).Decode(&conf)
		if err != nil {
			panic(err)
		}
	}

	log.Println("config loaded")

	Repos = conf.Repo
	ListenAddr = conf.ListenAddr
}
