package ioutil

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ReadAll(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		fmt.Println("oh my god")
	}
	return buf.Bytes(), err
}
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}
func TempDir(dir, prefix string) (name string, err error) {
	return ioutil.TempDir(dir, prefix)
}
