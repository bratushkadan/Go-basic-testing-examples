package util

import "testing"

func TestPersons(t *testing.T) {
	p1 := Person{"Dan", "Bratushka", "O."}
	p2 := Person{"Andrew", "Skochek", "P."}

	if expected, actual := "Dan O. Bratushka", p1.String(); expected != actual {
		t.Errorf("Name string conversion expected to be %s, got %s", expected, actual)
	}
	if expected, actual := "Andrew P. Skochek", p2.String(); expected != actual {
		t.Errorf("Name string conversion expected to be %s, got %s", expected, actual)
	}
}

func TestPersonWithEmptyFathersName(t *testing.T) {
	p1 := &Person{"Dan", "Bratushka", ""}

	if expected, actual := "Dan Bratushka", p1.String(); expected != actual {
		t.Errorf("Name string conversion expected to be %s, got %s", expected, actual)
	}
}
