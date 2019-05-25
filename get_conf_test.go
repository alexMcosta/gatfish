package main

import (
	"testing"
)

func TestGetConf(t *testing.T) {
	var c Configure
	got := c.getConf()
	want := &c

	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}
