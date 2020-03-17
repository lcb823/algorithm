package stack

import (
	"testing"
)

func TestBrowser(t *testing.T) {
	browser := NewBrowser()
	browser.Open("http://google.com")
	browser.Open("http://bing.com")
	browser.Open("http://baidu.com")

	if browser.CanForward() {
		t.Fatal("check CanForward error when false")
	}

	if !browser.CanBack() {
		t.Fatal("check CanBack error")
	}
	browser.Back()

	if !browser.CanForward() {
		t.Fatal("check CanForward error when true")
	}

	browser.Back()

	browser.Back()

	browser.Forward()

	browser.Forward()

}