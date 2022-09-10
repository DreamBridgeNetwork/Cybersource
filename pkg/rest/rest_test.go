package rest

import "testing"

func TestLoadCybersourceConfiguration(t *testing.T) {

	err := LoadCybersourceConfiguration()

	if err != nil {
		t.Error("rest.TestLoadCybersourceConfiguration = ", err.Error())
	}
}
