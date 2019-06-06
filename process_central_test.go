package main

import "testing"

type Tags struct {
	ALL map[string]string
	EC2 map[string]string
	EBS map[string]string
}

func TestProcessCentral(t *testing.T) {

	tags := Tags{
		ALL: map[string]string{"Name": "DIRE"},
		EC2: map[string]string{"Owner": "DIRE"},
		EBS: map[string]string{"Team": "DIRE"},
	}

	tagPointer := &tags

	got := processCentral(tagPointer)
	want := "Process Central Complete"

	if got != want {
		t.Errorf("wanted calls %v got %v", want, got)
	}
}
