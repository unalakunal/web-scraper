package main

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

func Test_searchInHTML(t *testing.T) {
	f, err := ioutil.ReadFile("test.html")
	if err != nil {
		panic(err)
	}
	r := bytes.NewReader(f)
	returned := searchInHTML(r)
	wantLinks := []string{"http://example.com/1", "http://example.com/2", "http://example.com/3", "http://example.com/4"}
	if !reflect.DeepEqual(returned, wantLinks) {
		t.Fail()
	}
}
