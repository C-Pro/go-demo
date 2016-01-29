package mypack

import "testing"

func TestGetName(t *testing.T) {
	name := GetName()
	if name != "Мумий Тролль" {
		t.Errorf("GetName() returned unexpected result: %s\n", name)
	}
}
