package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func SaveFile(fileName string, content []string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, line := range content {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func ListFile(dir string) ([]string, error) {

	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		fileNames = append(fileNames, fmt.Sprintf("%s%s", dir, f.Name()))
	}
	return fileNames, err
}
