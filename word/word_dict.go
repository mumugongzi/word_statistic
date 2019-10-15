package word

import (
	"fmt"
	"github.com/mumugongzi/word_statistic/util"
	"sort"
)

type WordDict struct {
	dict map[string]*Word
}

func NewWordDict() *WordDict {
	return &WordDict{
		dict: map[string]*Word{},
	}
}

func (d *WordDict) CountWord(word string) {
	w, ok := d.dict[word]
	if !ok {
		w = NewWord(word)
		d.dict[word] = w
	}
	w.Count()
}

func (d *WordDict) GetSortedWordList() WordList {
	var wordList WordList
	for _, v := range d.dict {
		wordList = append(wordList, v)
	}
	sort.Sort(wordList)
	return wordList
}

func (d *WordDict) Print() {
	wordList := d.GetSortedWordList()
	for _, w := range wordList {
		fmt.Printf("%s\n", w.ToString())
	}
}

func (d *WordDict) Save(fileName string) error {
	wordList := d.GetSortedWordList()
	var strArr []string
	for _, w := range wordList {
		strArr = append(strArr, w.ToString())
	}
	return util.SaveFile(fileName, strArr)
}
