package propsReader

import (
	"bytes"
	"errors"
	"github.com/guoapeng/props/utils"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

type AppConfigProperties interface {
	Get(key string) string
}

type appConfigProperties struct {
	props map[string]string
}

func (appConf *appConfigProperties) Get(key string) string {
	return appConf.props[key]
}

type AppConfigFactory struct {
	propertyFile string
	SystemDir    string
	HomeDir      string
	OsUtils      utils.OsUtils
	BufioUtils   utils.BufioUtils
}

func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

func NewFactory(appName string, propertyFile string) *AppConfigFactory {
	if homedir, err := Home(); err == nil {
		return &AppConfigFactory{OsUtils: utils.NewOsUtils(), BufioUtils: utils.NewBufioUtils(), SystemDir: "/etc/" + appName + "/", HomeDir: homedir + "/." + appName + "/", propertyFile: propertyFile}
	} else {
		log.Fatal("can't find home directory")
	}
	return nil
}

func (factory *AppConfigFactory) New() (AppConfigProperties, error) {
	var appConfigFile string
	appConfigFile = factory.propertyFile
	if len(appConfigFile) == 0 {
		log.Fatal("mandatory property file is not set")
	}

	systemProps, _ := factory.ReadPropertiesFile(factory.SystemDir + appConfigFile)
	homeProps1, _ := factory.ReadPropertiesFile(factory.HomeDir + appConfigFile)

	for k, v := range homeProps1 {
		systemProps[k] = v
	}
	return &appConfigProperties{props: systemProps}, nil
}

func (factory *AppConfigFactory) ReadPropertiesFile(filename string) (map[string]string, error) {
	config := map[string]string{}
	if len(filename) == 0 {
		return config, nil
	}

	if exists, err := factory.OsUtils.PathExists(filename); !exists || err != nil {
		return config, err
	}

	file, err := factory.OsUtils.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := factory.BufioUtils.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		preProcessedLine := strings.TrimSpace(string(line))
		if strings.HasPrefix(preProcessedLine, "source") {
			position := strings.Index(preProcessedLine, "source")
			if tempProps, err := factory.ReadPropertiesFile(strings.TrimSpace(preProcessedLine[position+6:])); err == nil {
				for k, v := range tempProps {
					config[k] = v
				}
			}
		}
		if strings.HasPrefix(preProcessedLine, "#") {
			continue
		}
		if equal := strings.Index(preProcessedLine, "="); equal >= 0 {
			if key := strings.TrimSpace(preProcessedLine[:equal]); len(key) > 0 {
				value := ""
				if len(preProcessedLine) > equal {
					value = strings.TrimSpace(preProcessedLine[equal+1:])
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
