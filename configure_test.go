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

	t.Run("get one value in tag map", func(t *testing.T) {
		var c Configure
		c.GetConf()
		got := &c.EC2
		want := assertWant(t, "Name", "DIRE")

		assertMaps(t, got, want)
	})

	t.Run("get multiple in tag map", func(t *testing.T) {

		var c Configure
		c.GetConf()
		wantMap := make(map[string]string)
		wantMap["Owner"] = "REQUIRED"
		wantMap["Name"] = "DIRE"

		got := &c.EBS
		want := &wantMap

		assertMaps(t, got, want)
	})

}

func assertWant(t *testing.T, tag, value string) *map[string]string {
	t.Helper()

	wantMap := make(map[string]string)
	wantMap[tag] = value
	return &wantMap
}

func assertMaps(t *testing.T, got, want *map[string]string) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted calls %v got %v", want, got)
	}
}
