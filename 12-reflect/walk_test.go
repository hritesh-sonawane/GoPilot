package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"Nested fields",
			Person{
				"Tanjiro",
				Profile{14, "Japan"},
			},
			[]string{"Tanjiro", "Japan"},
		},
		{
			"Pointers to things",
			&Person{
				"Tanjiro",
				Profile{14, "Japan"},
			},
			[]string{"Tanjiro", "Japan"},
		},
		{
			"Slices",
			[]Profile{
				{14, "Japan"},
				{17, "China"},
			},
			[]string{"Japan", "China"},
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
