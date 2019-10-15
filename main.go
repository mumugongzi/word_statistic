package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/mumugongzi/word_statistic/doc"
	"github.com/mumugongzi/word_statistic/logs"
	"github.com/mumugongzi/word_statistic/util"
	"github.com/mumugongzi/word_statistic/word"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	dir *string

	dict      = word.NewWordDict()
	splitList = []string{" ", "\n", ":", "'", "\"", ",", ".", "!", "(", ")", "?", "、", "“", "’", "_", "-", "。", "（", "；", "；", "＄", "”", " ", "—", "，", "！", "）", "．", "：", "‘"}
	trimList  = []string{":", "'", "\"", ",", ".", "'s", "!", "(", ")", "?", "、", "“", "’", "_", "-", "。", "—", "√", "–", "＼", "…"}

	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	if !validParam() {
		return
	}

	fileNames, err := util.ListFile(*dir)
	if err != nil {
		logs.Error("list file err: %v", err)
	}

	fileNames = filterNames(fileNames)

	for _, f := range fileNames {
		logs.Info("start static file: %s", f)
		textList, err := doc.ReadText(f)
		if err != nil {
			logs.Error("read file: %s err: %v", f, err)
		}

		for _, text := range textList {
			wordArr := util.Split(text, splitList)
			for _, word := range wordArr {
				word = strings.ToLower(word)
				word = util.Trim(word, trimList, true)
				if len(word) <= 1 {
					continue
				}

				if filterWord(word) {
					continue
				}
				dict.CountWord(word)
			}
		}
	}

	var resFile string
	if strings.HasSuffix(*dir, "/") {
		resFile = fmt.Sprintf("%s%s", *dir, "词频统计结果.txt")
	} else {
		resFile = fmt.Sprintf("%s/%s", *dir, "词频统计结果.txt")
	}
	dict.Print()
	err = dict.Save(resFile)
	if err != nil {
		logs.Error("save dict to file %s err: %v", resFile, err)
	}
}

func validParam() bool {

	dir = flag.String("d", "", "设置文件路径")
	flag.Parse()

	if *dir == "" {
		//fmt.Println("please provider a dir")
		//flag.Usage()
		//return false
		str, err := GetExecPath()
		if err != nil {
			panic(err)
		}
		*dir = str
		logs.Info("cur dir: %s", *dir)
	}
	return true
}

func filterNames(fileNames []string) []string {

	var res []string
	suffixList := []string{"doc", "docx"}
	for _, f := range fileNames {
		reserve := false
		for _, suffix := range suffixList {
			if strings.HasSuffix(f, suffix) {
				reserve = true
				break
			}
		}

		if reserve {
			res = append(res, f)
		}
	}
	return res
}

func filterWord(word string) bool {
	if len(word) <= 0 {
		return false
	}
	for i := 0; i < len(alphabet); i++ {
		if strings.HasPrefix(word, alphabet[i:i+1]) {
			return false
		}
	}
	return true
}

func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
