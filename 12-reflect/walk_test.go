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
		{
			"Arrays",
			[2]Profile{
				{14, "Japan"},
				{17, "China"},
			},
			[]string{"Japan", "China"},
		},
	}
	// {
	// 	"Maps",
	// 	map[string]string{
	// 		"Foo": "Bar",
	// 		"Baz": "Boz",
	// 	},
	// 	[]string{"Bar", "Boz"},
	// },
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{14, "Japan"}
			aChannel <- Profile{17, "China"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Japan", "China"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v | want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{14, "Japan"}, Profile{17, "China"}
		}

		var got []string
		want := []string{"Japan", "China"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v | want %v", got, want)
		}
	})

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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
