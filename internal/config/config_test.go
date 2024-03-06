package config

import "testing"

func TestLoadTemplates(t *testing.T) {
	ts, err := LoadTemplates("./testdata/.git-czrc")
	if err != nil {
		t.Log(err)
	}
	t.Log(ts)
}
