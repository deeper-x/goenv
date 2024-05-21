package goenv

import (
	"testing"
)

func TestFileDump(t *testing.T) {
	got, err := FileDump("./assets/.env")
	if err != nil {
		t.Error(err)
	}

	expLines := []string{`USER=demo_user`, `PASSWORD=qwerty`}

	c := 0
	for i := range expLines {
		if got[c] != expLines[i] {
			t.Errorf("%s != %s", got[c], expLines[i])
		}
		c++
	}
}

func TestEmptyInputParam(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Error(err)
	}
}

func TestNonEmptyInputParam(t *testing.T) {
	_, err := New("./assets/.env")
	if err != nil {
		t.Error(err)
	}
}

func TestBuildKV(t *testing.T) {
	res, err := buildKV([]string{"par1=one", "par2=two"})
	if err != nil {
		t.Error(err)
	}

	exp1 := "one"
	if res["par1"] != exp1 {
		t.Errorf("expected %s, got %s", "one", exp1)
	}

	exp2 := "two"
	if res["par2"] != exp2 {
		t.Errorf("expected %s, got %s", "two", exp2)
	}
}

func TestGetOK(t *testing.T) {
	efo, err := New()
	if err != nil {
		t.Error(err)
	}

	v, err := efo.Get("USER")
	if err != nil {
		t.Error(err)
	}

	expected := "demo_user"
	if v != "demo_user" {
		t.Errorf("expected %s, got %s", expected, v)
	}
}

func TestCheckRowsFormatOK(t *testing.T) {
	rows := []string{
		"foo=1",
		"bar_1=2",
		"baz-2=3",
	}

	for i := range rows {
		ok, err := checkRowFormat(rows[i])
		if err != nil {
			t.Error(err)
		}

		if !ok {
			t.Error("input entries are expected to be format compliant")
		}
	}

}

func TestCheckRowsFormatKO(t *testing.T) {
	rows := []string{
		"doo*1",
		"wrong_record",
		"baz 2",
	}

	for _, v := range rows {
		ok, err := checkRowFormat(v)
		if err != nil {
			t.Error(err)
		}

		if ok {
			t.Errorf("input entry %s is expected not to be format compliant", v)
		}
	}

}

func TestGetKO(t *testing.T) {
	efo, err := New()
	if err != nil {
		t.Error(err)
	}

	_, err = efo.Get("NON_EXISTING")
	if err == nil {
		t.Error("Get method should return an error, because no existing key fetched")
	}
}
