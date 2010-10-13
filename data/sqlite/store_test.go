package store

import "testing"
import "os"
import "mudkip/lib"

func TestObject(t *testing.T) {
	var err os.Error

	ds := New()
	if err = ds.Open(map[string]string{"file": "test.db"}); err != nil {
		t.Errorf("Open: %v", err)
		return
	}

	defer ds.Close()

	if err = ds.Initialize(); err != nil {
		t.Errorf("Initialize: %v", err)
		return
	}

	var a, b lib.Object

	a = lib.NewZone()
	a.SetId(1)
	a.SetName("lobby")
	a.SetDescription("Entrance to our great domain")

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


func TestUser(t *testing.T) {
	var err os.Error

	ds := New()
	if err = ds.Open(map[string]string{"file": "test.db"}); err != nil {
		t.Errorf("Open: %v", err)
		return
	}

	defer ds.Close()

	if err = ds.Initialize(); err != nil {
		t.Errorf("Initialize: %v", err)
		return
	}

	var a, b *lib.UserInfo

	a = new(lib.UserInfo)
	a.Id = 1
	a.Name = "bob"
	a.Password = "1234"
	a.Registered = 0
	a.Zone = 1

	if err = ds.SetUser(a); err != nil {
		t.Errorf("SetUser: %v", err)
		return
	}

	if b, err = ds.GetUser(a.Id); err != nil {
		t.Errorf("GetUser: %v", err)
		return
	}

	if b == nil || a.Name != b.Name {
		t.Errorf("GetUser: users do not match:\n- %#v\n- %#v", a, b)
		return
	}

	if b, err = ds.GetUserByName(a.Name); err != nil {
		t.Errorf("GetUserByName: %v", err)
		return
	}

	if b == nil || a.Name != b.Name {
		t.Errorf("GetUserByName: users do not match:\n- %#v\n- %#v", a, b)
		return
	}
}
