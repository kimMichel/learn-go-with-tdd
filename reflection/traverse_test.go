package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func TestTraverse(t *testing.T) {
	cases := []struct {
		Name string
		Input interface{}
		ExpectedRequests []string
	}{
		{
			"Struct with string field",
			struct{
				Name string
			}{"Michel"},
			[]string{"Michel"},
		},
		{
			"Struct with two string field",
			struct{
				Name string
				City string
			}{"Michel", "London"},
			[]string{"Michel", "London"},
		},
		{
			"Struct without string field",
			struct{
				Name string
				Age int
			}{"Michel", 25},
			[]string{"Michel"},
		},
		{
			"Nested fields",
			Person{
				"Michel",
				Profile{25, "London"},
			},
			[]string{"Michel", "London"},
		},
		{
			"Pointers",
			&Person{
				"Michel",
				Profile{25, "London"},
			},
			[]string{"Michel", "London"},
		},
		{
			"Slices",
			[]Profile{
				{25, "London"},
				{27, "Seoul"},
			},
			[]string{"London", "Seoul"},
		},
		{
			"Arrays",
			[2]Profile{
				{25, "London"},
				{27, "Seoul"},
			},
			[]string{"London", "Seoul"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string
			traverse(test.Input, func(input string) {
				result = append(result, input)
			})

			if !reflect.DeepEqual(result, test.ExpectedRequests) {
				t.Errorf("result: %v, expect: %v", result, test.ExpectedRequests)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		mapA := map[string]string{
			"Foo": "Bar",
        	"Baz": "Boz",
		}

		var result []string
		traverse(mapA, func(input string) {
			result = append(result, input)
		})

		checkIfHas(t, result, "Bar")
		checkIfHas(t, result, "Boz")
	})
}

func checkIfHas(t *testing.T, haystack []string, needle string) {
	has := false
	for _, x := range haystack {
		if x == needle {
			has = true
		}
	}

	if !has {
		t.Errorf("expected to has %v in %v, but does not has", needle, haystack)
	}
}