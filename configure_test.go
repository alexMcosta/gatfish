package main

import (
	"reflect"
	"testing"
)

func TestGetConf(t *testing.T) {

	t.Run("get pointer to a configure object", func(t *testing.T) {
		var c Configure
		got := c.GetConf()
		want := &c

		if want != got {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("get EC2 tag map", func(t *testing.T) {
		var c Configure
		wantMap := make(map[string]string)
		wantMap["Name"] = "DIRE"
		c.GetConf()
		got := &c.EC2
		want := &wantMap

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
