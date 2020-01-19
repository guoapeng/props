package utils

import (
	"bufio"
	"io"
)

type BufioUtils interface {
	NewScanner(r io.Reader) *bufio.Scanner
}

func NewBufioUtils() BufioUtils {
	return &bufioUtils{}
}

type bufioUtils struct {
}

func (b *bufioUtils) NewScanner(r io.Reader) *bufio.Scanner {
	return bufio.NewScanner(r)
}
