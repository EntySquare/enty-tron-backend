package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Conf struct {
	Postgres struct {
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		User       string `yaml:"user"`
		Password   string `yaml:"password"`
		Dbname     string `yaml:"dbname"`
		ExecSchema bool   `yaml:"execSchema"`
	} `yaml:"postgres"`
}

func (m *Conf) GetConf() *Conf {
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("base path error")
	}
	index := strings.Index(basePath, "enty-tron-backend")
	fileName := filepath.Join(basePath[0:index+17], "conf.yaml")
	fmt.Println("config:", fileName)
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("load conf error")
	}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		fmt.Println(err.Error())
	}
	return m
}
