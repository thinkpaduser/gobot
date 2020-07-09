package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type MsgDB map[string][]string

type Config struct {
	Conf struct {
		Token  string `yaml:"token"`
		Debug  bool   `yaml:"debug"`
		Target string `yaml:"target"`
	}
}

func NewConf(filename string) (c *Config, err error) {
	c = &Config{}

	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return
	}
	defer file.Close()

	var bytes []byte
	if bytes, err = ioutil.ReadAll(file); err != nil {
		err = fmt.Errorf("Can't read file %s: %s", filename, err.Error())
		return
	}
	if err = yaml.Unmarshal(bytes, c); err != nil {
		err = fmt.Errorf("Can't parse file %s: %s", filename, err.Error())
		return
	}
	return
}

func NewMsgDB(filename string) (m map[string][]string, err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return
	}
	defer file.Close()
	var bytes []byte
	if bytes, err = ioutil.ReadAll(file); err != nil {
		err = fmt.Errorf("Can't read file %s: %s", filename, err.Error())
		return
	}

	if err = yaml.Unmarshal(bytes, &m); err != nil {
		err = fmt.Errorf("Can't parse file %s: %s", filename, err.Error())
		return
	}
	return
}
