package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	conf struct {
		Token string `yaml:"token"`
		Debug bool   `yaml:"debug"`
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
