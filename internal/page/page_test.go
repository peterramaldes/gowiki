package page

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	f := "foo"
	e := ".txt"
	filename := fmt.Sprintf("%s%s", f, e)
	p := &Page{Title: f, Body: []byte("bar")}
	defer os.Remove(filename)

	err := p.Save()
	if err != nil {
		t.Fatalf("error trying to save page: %s", err)
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		t.Errorf("file %s doesn't exists", filename)
	}
}

func TestNew(t *testing.T) {
	title := "x"
	filename := fmt.Sprintf("%s.txt", title)
	f, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}

	body := "y"
	err = os.WriteFile(f.Name(), []byte(body), 0600)
	if err != nil {
		t.Fatal(err)
	}

	p, err := New(title)
	if err != nil {
		t.Fatal(err)
	}

	expected := &Page{Title: title, Body: []byte(body)}
	if !deepEquals(expected, p) {
		t.Errorf("expected %s, but return %s", expected, p)
	}

}

func deepEquals(p1 *Page, p2 *Page) bool {
	if p1.Title != p2.Title {
		return false
	}
	if !reflect.DeepEqual(p1.Body, p2.Body) {
		return false
	}
	return true
}
