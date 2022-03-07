package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOK = `1
2
3
3
4
5`

var testOkResult = `1
2
3
4
5`

func testOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOK))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test for OK Failed")
	}
	if out.String() != testOkResult {
		t.Errorf("test for OK Failed - results not watch\n %v %v", out.String(), testOkResult)
	}
}

var testFail = `1
2
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("test for OK Failed - error: %v", err)
	}
}
