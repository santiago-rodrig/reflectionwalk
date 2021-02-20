package reflectionwalk

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}

type Person struct {
	Name string
	Profile
}

func TestWalk(t *testing.T) {
	t.Run("Not maps", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"Struct with one string field",
				struct {
					Name string
				}{"Santiago"},
				[]string{"Santiago"},
			},
			{
				"Struct with two string fields",
				struct {
					Name string
					City string
				}{"Santiago", "Venezuela"},
				[]string{"Santiago", "Venezuela"},
			},
			{
				"Struct with non string field",
				struct {
					Name string
					Age  int
				}{"Santiago", 24},
				[]string{"Santiago"},
			},
			{
				"Nested fields",
				Person{
					"Santiago",
					Profile{24, "Venezuela"},
				},
				[]string{"Santiago", "Venezuela"},
			},
			{
				"Pointers to things",
				&Person{
					"Santiago",
					Profile{24, "Venezuela"},
				},
				[]string{"Santiago", "Venezuela"},
			},
			{
				"Slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
			{
				"Arrays",
				[2]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string

				Walk(test.Input, func(input string) {
					got = append(got, input)
				})

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, test.ExpectedCalls)
				}
			})
		}
	})

	t.Run("With maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
