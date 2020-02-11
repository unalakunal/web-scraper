package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_searchInHTML(t *testing.T) {
	f, err := os.Open("test.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var r io.ReadCloser = f
	returned := searchInHTML(r)
	wantLinks := []string{"http://example.com/1", "http://example.com/2", "http://example.com/3", "http://example.com/4"}
	if !reflect.DeepEqual(returned, wantLinks) {
		t.Fail()
	}
}
