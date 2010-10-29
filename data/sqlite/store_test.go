package store

import "testing"

func Test(t *testing.T) {
	ds := New()
	if err := ds.Open(map[string]string{"file": "test.db"}); err != nil {
		t.Error(err.String())
		return
	}

	defer ds.Close()

	if err := ds.Initialize(nil); err != nil {
		t.Error(err.String())
	}

}
