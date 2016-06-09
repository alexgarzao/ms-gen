package main

import "testing"

type user struct {
	Id   int
	Name string `valid:"Required"`
	Age  int    `valid:"Required"`
}

func TestValid(t *testing.T) {
	cases := []struct {
		in             user
		expectedState  bool
		expectedResult string
	}{
		{user{Name: "alex", Age: 41}, true, ""},
		{user{Name: "alex"}, false, "Age"},
		{user{Id: 111, Age: 40}, false, "Name"},
		{user{Id: 111}, false, "Name,Age"},
	}
	for _, c := range cases {
		got, missed := CheckRequiredFields(c.in)
		if got != c.expectedState {
			t.Errorf("CheckRequiredFields(%v) == %v, want %v", c.in, got, c.expectedState)
		}
		if missed != c.expectedResult {
			t.Errorf("CheckRequiredFields(%v) == '%v', want '%v'", c.in, missed, c.expectedResult)
		}
	}
}
