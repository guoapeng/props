package propsReader_test

import (
	"bufio"
	"github.com/golang/mock/gomock"
	"github.com/guoapeng/props"
	"github.com/guoapeng/props/mocks"
	"github.com/stretchr/testify/suite"
	"os"
	"strings"
	"testing"
)

func TestPropsReaderSuite(t *testing.T) {
	suite.Run(t, new(PropsReaderSuite))
}

type PropsReaderSuite struct {
	suite.Suite
	test         *testing.T
	osUtils      *mocks.MockOsUtils
	bufioUtils   *mocks.MockBufioUtils
	factory      *propsReader.AppConfigFactory
	SystemFolder string
	HomeDir      string
}

func (s *PropsReaderSuite) T() *testing.T {
	return s.test
}

func (s *PropsReaderSuite) SetT(t *testing.T) {
	s.test = t
	mockCtrl := gomock.NewController(s.T())
	s.osUtils = mocks.NewMockOsUtils(mockCtrl)
	s.bufioUtils = mocks.NewMockBufioUtils(mockCtrl)
	s.SystemFolder = "/etc/APP/"
	homedir, _ := propsReader.Home()
	s.HomeDir = homedir + "/.APP/"
	s.factory = propsReader.NewFactory("APP", "config.properties")
	s.factory.OsUtils = s.osUtils
	s.factory.BufioUtils = s.bufioUtils
	defer mockCtrl.Finish()
}

func (s *PropsReaderSuite) TestReadingConfigFromSystemFolder() {
	s.osUtils.EXPECT().Getenv("APP_CONFIG").Return("config.properties")
	fileInSystemDir := &os.File{}
	fileInHome := &os.File{}
	s.osUtils.EXPECT().Open(s.SystemFolder+"config.properties").Return(fileInSystemDir, nil)
	s.osUtils.EXPECT().Open(s.HomeDir+"config.properties").Return(fileInHome, nil)
	s.osUtils.EXPECT().PathExists(s.SystemFolder+"config.properties").Return(true, nil)
	s.osUtils.EXPECT().PathExists(s.HomeDir+"config.properties").Return(true, nil)
	buf1 := strings.NewReader("key=value\n key2 = value2 ")
	buf2 := strings.NewReader("key=valueFromHome\n key3 = value3 ")
	firstScanner := bufio.NewScanner(buf1)
	secondScanner := bufio.NewScanner(buf2)
	s.bufioUtils.EXPECT().NewScanner(fileInSystemDir).Return(firstScanner)
	s.bufioUtils.EXPECT().NewScanner(fileInHome).Return(secondScanner)
	if appConf, err := s.factory.New(); err != nil {
		s.T().Errorf("reading config file error")
	} else {
		if appConf.Get("key") != "valueFromHome" {
			s.T().Errorf("paring config file error")
		}
		if appConf.Get("key2") != "value2" {
			s.T().Errorf("paring config file error, %s is missing", "key2")
		}
		if appConf.Get("key3") != "value3" {
			s.T().Errorf("paring config file error, %s is missing", "key3")
		}
	}
}

func (s *PropsReaderSuite) TestReadingConfigFromSystemFolderAndReplacePlaceHoldersWithEnvVariables() {
	s.osUtils.EXPECT().Getenv("APP_CONFIG").Return("config.properties")
	fileInSystemDir := &os.File{}
	fileInHome := &os.File{}
	s.osUtils.EXPECT().Open(s.SystemFolder+"config.properties").Return(fileInSystemDir, nil)
	s.osUtils.EXPECT().Open(s.HomeDir+"config.properties").Return(fileInHome, nil)
	s.osUtils.EXPECT().PathExists(s.SystemFolder+"config.properties").Return(true, nil)
	s.osUtils.EXPECT().PathExists(s.HomeDir+"config.properties").Return(true, nil)
	s.osUtils.EXPECT().Getenv("Server_addr").Return("192.169.200.2")
	s.osUtils.EXPECT().Getenv("Server_port").Return("8888")
	s.osUtils.EXPECT().Getenv("JAVA_HOME").Return("xx")
	buf1 := strings.NewReader("key=value\n key2 = value2 \n key4=http://${Server_addr}:${Server_port}")
	buf2 := strings.NewReader("key=valueFromHome\n key3 = value3 ")
	firstScanner := bufio.NewScanner(buf1)
	secondScanner := bufio.NewScanner(buf2)
	s.bufioUtils.EXPECT().NewScanner(fileInSystemDir).Return(firstScanner)
	s.bufioUtils.EXPECT().NewScanner(fileInHome).Return(secondScanner)
	if appConf, err := s.factory.New(); err != nil {
		s.T().Errorf("reading config file error")
	} else {
		if appConf.Get("key") != "valueFromHome" {
			s.T().Errorf("paring config file error")
		}
		if appConf.Get("key2") != "value2" {
			s.T().Errorf("paring config file error, %s is missing", "key2")
		}
		if appConf.Get("key3") != "value3" {
			s.T().Errorf("paring config file error, %s is missing", "key3")
		}

		if appConf.Get("key4") == "http://${Server_addr}:${Server_port}" {
			s.T().Errorf("paring config file error, %s is expected to be replaced but not", "key4")
		}

		if appConf.Get("key4") != "http://192.169.200.2:8888" {
			s.T().Errorf("paring config file error, %s is missing", "key4")
		}
	}
}

func (s *PropsReaderSuite) TestSourceFunction() {

	s.osUtils.EXPECT().Getenv("APP_CONFIG").Return("config.properties")
	fileInSystemDir := &os.File{}
	fileInHome := &os.File{}
	sourceInlineFile := &os.File{}

	firstScanner := bufio.NewScanner(strings.NewReader("key=value\n key2 = value2 \n source config2.properties"))
	secondScanner := bufio.NewScanner(strings.NewReader("key=valueFromHome \n key3 = value3 "))
	thirdScanner := bufio.NewScanner(strings.NewReader("key4=value4 \n key3 = value4 "))

	s.bufioUtils.EXPECT().NewScanner(fileInSystemDir).Return(firstScanner)
	s.bufioUtils.EXPECT().NewScanner(sourceInlineFile).Return(thirdScanner)
	s.bufioUtils.EXPECT().NewScanner(fileInHome).Return(secondScanner)
	s.osUtils.EXPECT().Open(s.SystemFolder+"config.properties").Return(fileInSystemDir, nil)
	s.osUtils.EXPECT().Open(s.HomeDir+"config.properties").Return(fileInHome, nil)
	s.osUtils.EXPECT().Open("config2.properties").Return(sourceInlineFile, nil)
	s.osUtils.EXPECT().PathExists(s.SystemFolder+"config.properties").Return(true, nil)
	s.osUtils.EXPECT().PathExists(s.HomeDir+"config.properties").Return(true, nil)
	s.osUtils.EXPECT().PathExists("config2.properties").Return(true, nil)

	if appConf, err := s.factory.New(); err != nil {
		s.T().Errorf("reading config file error")
	} else {
		if appConf.Get("key") != "valueFromHome" {
			s.T().Errorf("paring config file error")
		}
		if appConf.Get("key2") != "value2" {
			s.T().Errorf("paring config file error, %s is missing", "key2")
		}
		if appConf.Get("key3") != "value3" {
			s.T().Errorf("paring config file error, %s is missing", "key3")
		}
		if appConf.Get("key4") != "value4" {
			s.T().Errorf("paring config file error, %s is missing", "key4")
		}
	}
}

func (s *PropsReaderSuite) TestDoNothingWhenSourceFileDoesntExist() {

	s.osUtils.EXPECT().Getenv("APP_CONFIG").Return("config.properties")
	fileInSystemDir := &os.File{}
	fileInHome := &os.File{}
	sourceInlineFile := &os.File{}

	firstScanner := bufio.NewScanner(strings.NewReader("key=value\n key2 = value2 \n source config2.properties \n source config3.properties "))
	secondScanner := bufio.NewScanner(strings.NewReader("key=valueFromHome \n key3 = value3 "))
	thirdScanner := bufio.NewScanner(strings.NewReader("key4=value4 \n key3 = value4 "))

	s.bufioUtils.EXPECT().NewScanner(fileInSystemDir).Return(firstScanner)
	s.bufioUtils.EXPECT().NewScanner(fileInHome).Return(secondScanner)
	s.bufioUtils.EXPECT().NewScanner(sourceInlineFile).Return(thirdScanner)
	s.osUtils.EXPECT().Open(s.SystemFolder+"config.properties").Return(fileInSystemDir, nil)
	s.osUtils.EXPECT().Open(s.HomeDir+"config.properties").Return(fileInHome, nil)
	s.osUtils.EXPECT().Open("config2.properties").Return(sourceInlineFile, nil)
	s.osUtils.EXPECT().PathExists(s.SystemFolder+"config.properties").Return(true, nil)
	s.osUtils.EXPECT().PathExists(s.HomeDir+"config.properties").Return(true, nil)
	s.osUtils.EXPECT().PathExists("config2.properties").Return(false, nil)
	s.osUtils.EXPECT().PathExists("config3.properties").Return(false, nil)
	if appConf, err := s.factory.New(); err != nil {
		s.T().Errorf("reading config file error")
	} else {
		if appConf.Get("key") != "valueFromHome" {
			s.T().Errorf("paring config file error")
		}
		if appConf.Get("key2") != "value2" {
			s.T().Errorf("paring config file error, %s is missing", "key2")
		}
		if appConf.Get("key3") != "value3" {
			s.T().Errorf("paring config file error, %s is missing", "key3")
		}
		if appConf.Get("key4") == "value4" {
			s.T().Errorf("paring config file error, %s shouldn't be loaded", "key4")
		}
	}
}
