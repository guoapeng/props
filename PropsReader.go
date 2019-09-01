package propsReader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const configFile = "config.properties"

type AppConfigProperties struct{
	props map[string]string
}

func (appConf *AppConfigProperties) Get(key string) string {
	return appConf.props[key]
}

func New() (*AppConfigProperties, error) {
	if props1, err := ReadPropertiesFile(configFile); err != nil {
		return nil, err
	} else {
		return &AppConfigProperties{props: props1}, nil
	}
}

func ReadPropertiesFile(filename string) (map[string]string, error) {
	config := map[string]string{}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return config, nil
}
