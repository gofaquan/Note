package code_test

import (
	"fmt"
	code "github.com/gofaquan"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	F1 string
	F2 int
}

func TestGob(t *testing.T) {
	should := assert.New(t)
	gobBytes, err := code.GobEncode(&TestStruct{F1: "test_f1", F2: 10})
	if should.NoError(err) {
		fmt.Println(gobBytes)
	}

	obj := TestStruct{}
	err = code.GobDecode(gobBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
