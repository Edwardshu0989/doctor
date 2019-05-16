package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

type AppConfig struct {
	// basic info

	Db struct {
		Server   string
		Dbname   string
		User     string
		Password string
	}
	Rediss struct {
		MaxIdle   int
		MaxActive int
		Host      string
		Dg        string
	}
}
type Config struct {
	cfg  *AppConfig
	file string
}

func (this *Config) SetFile(file string) error {
	this.file = file
	return nil
}
func (this *Config) Load() error {

	f, err := os.Open(this.file)
	defer f.Close()

	buff, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("Config file read fail, err: %s", err)
	}

	var r AppConfig
	err = toml.Unmarshal(buff, &r)
	if err != nil {
		return fmt.Errorf("Config file Unmarshal fail, err: %s", err)
	}
	this.cfg = &r

	return nil
}

func (this *Config) Get() *AppConfig {
	return this.cfg
}

var Default *Config = &Config{}

func SetFile(file string) error {
	return Default.SetFile(file)
}

func SetFileAndLoad(file string) error {
	Default.SetFile(file)
	return Default.Load()
}

func Get() *AppConfig {
	return Default.Get()
}
