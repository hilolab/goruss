package util

import (
	//"fmt"
	"bytes"
	"io"
	"os"
	"strings"
)

//分割路径
func PathSplit(path string) []string {
	path = strings.Trim(path, "/")
	segment := strings.Split(path, "/")
	return segment
}

func Ucfirst(str string) string {
	b := []byte(str)
	b[0] += 32
	return string(b)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)

	return err == nil || os.IsExist(err)
}

func ReadAll(fd *os.File) ([]byte, error) {
	result := bytes.NewBuffer(nil)
	var buf [4096]byte
	for {
		n, err := fd.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}
