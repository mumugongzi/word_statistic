package util

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	str := "aaas;hjshjah]]gshagsha agsgah\nagsag\n"
	res := Split(str, []string{";","]]"," ", "\n"})
	t.Log(strings.Join(res, "\n"))
}


func TestTrim(t *testing.T) {
	str := "aaas;hjshjah]]gshagsha agsgah\nagsag\n"
	t.Log(Trim(str, []string{"\n", "g", "a", "s", ";"}, true))
}
