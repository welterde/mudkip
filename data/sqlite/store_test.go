package store

import "testing"
import "os"
import "mudkip/lib"

func Test(t *testing.T) {
	var ds lib.DataStore
	var err os.Error

	if ds = New(); ds == nil {
		t.Error("Server has been built without datastore support. Cannot continue.")
		return
	}

	if err = ds.Open(map[string]string{"file": "test.db"}); err != nil {
		t.Errorf("Open: %v", err)
		return
	}

	defer ds.Close()

	if !ds.Initialized() {
		if err = ds.Initialize(); err != nil {
			t.Errorf("Initialize: %v", err)
			return
		}
	}

	var a, b lib.Object

	a = lib.NewZone()
	a.SetId(1)
	a.SetName("lobby")
	a.SetDescription("Entrace to our great domain")

	if err = ds.SetObject(a); err != nil {
		t.Errorf("SetObject: %v", err)
		return
	}

	if b, err = ds.GetObject(1, lib.OTZone); err != nil {
		t.Errorf("GetObject: %v", err)
		return
	}

	if b == nil || a.Name() != b.Name() {
		t.Errorf("objects do not match:\n- %#v\n- %#v", a, b)
		return
	}
}
