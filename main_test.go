package reflectionwalk

import "testing"

func TestWalk(t *testing.T) {
	expected := "Santiago"
	var got []string
	x := struct {
		Name string
	}{expected}
	Walk(x, func(input string) {
		got = append(got, input)
	})
	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d, want %d", len(got), 1)
	}
}
