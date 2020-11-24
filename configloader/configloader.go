package configloader

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg interface{}, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg interface{}) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

// Load config and return it
func Load(configFilename string, cfg interface{}) {
	readFile(cfg, configFilename)
	readEnv(cfg)
}
