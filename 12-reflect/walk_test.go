package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Tanjiro"},
			[]string{"Tanjiro"},
		},
		{
			"Struct with 2 string fields",
			struct {
				Name string
				City string
			}{"Tanjiro", "Japan"},
			[]string{"Tanjiro", "Japan"},
		},
		{
			"Struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Tanjiro", 14},
			[]string{"Tanjiro"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q | want %q", got, test.ExpectedCalls)
			}
		})
	}

}
