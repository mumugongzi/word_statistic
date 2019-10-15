package util

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestListFile(t *testing.T) {
	files, err := ListFile("/Users/lisong/Downloads/")
	assert.Nil(t, err)

	t.Log(strings.Join(files, "\n"))
}
