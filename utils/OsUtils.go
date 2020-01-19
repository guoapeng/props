package utils

import "os"

type OsUtils interface {
	Getenv(key string) string
	Open(name string) (*os.File, error)
	PathExists(path string) (bool, error)
}

func NewOsUtils() OsUtils {
	return &osUtils{}
}

type osUtils struct {
}

func (u *osUtils) Getenv(key string) string {
	return os.Getenv(key)
}

func (u *osUtils) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (u *osUtils) PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
