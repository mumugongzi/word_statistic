package util

import "strings"

func Split(str string, sepList []string) []string {
	inputArr := []string{str}
	var outputErr []string
	for _, sep := range sepList {
		for _, input := range inputArr {
			if strings.TrimSpace(input) == "" {
				continue
			}
			outputErr = append(outputErr, strings.Split(input, sep)...)
		}
		inputArr = outputErr
		outputErr = []string{}
	}
	return inputArr
}

func Trim(str string, trimList []string, trimSpace bool) string {

	if trimSpace {
		str = strings.TrimSpace(str)
	}

	for {
		origin := len(str)
		for _, trimStr := range trimList {
			str = trimSingle(str, trimStr)
		}
		if origin == len(str) {
			break
		}
	}

	return str
}

func trimSingle(str string, trimStr string) string {

	trimStr = strings.TrimSpace(trimStr)
	if len(trimStr) <= 0 {
		return str
	}

	for {
		if !strings.HasPrefix(str, trimStr) {
			break
		}
		str = str[len(trimStr):]
	}

	for {
		if !strings.HasSuffix(str, trimStr) {
			break
		}
		str = str[:len(str)-len(trimStr)]
	}
	return str
}
