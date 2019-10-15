package word

import (
	"fmt"
	"strings"
)

type Word struct {
	word string
	cnt  int
}

func NewWord(word string) *Word {
	return &Word{
		word: word,
		cnt:  0,
	}
}

func (w *Word) Count() {
	w.cnt++
}

func (w *Word) ToString() string {
	return fmt.Sprintf("%s\t%d", w.word, w.cnt)
}

type WordList []*Word

func (ws WordList) Len() int {
	return len(ws)
}

func (ws WordList) Swap(i, j int) {
	ws[i], ws[j] = ws[j], ws[i]
}

// 先按词频倒序排, 后按字典序排
func (ws WordList) Less(i, j int) bool {
	if ws[i].cnt != ws[j].cnt {
		return ws[i].cnt > ws[j].cnt
	}

	ret := strings.Compare(ws[i].word, ws[j].word)
	if ret < 0 {
		return true
	} else {
		return false
	}
}
