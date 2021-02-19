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
}
