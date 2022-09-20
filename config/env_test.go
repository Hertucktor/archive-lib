package config

import "testing"

func TestGetConfig(t *testing.T) {
	c, err := GetConfig("config_test.yml")
	if err != nil {
		t.Error("There should not be an error whilst parsing the yaml file")
	}

	if c.DBName != "testDBName" {
		t.Fatalf("Expected: testDBName, got: %v ", c.DBName)
	}
}