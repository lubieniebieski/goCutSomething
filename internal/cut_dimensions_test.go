package cut_dimensions

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestGenerateIntegers(t *testing.T) {
	desiredInt := 100
	count := 5
	expected := []int{
		95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105,
	}

	got := generateIntegers(desiredInt, count)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("generateIntegers(%d) == %v, want %v", desiredInt, got, expected)
	}
}

func TestCuttableWithoutMargin(t *testing.T) {
	materialWidth := 1000
	diameter := 100
	spacing := 50
	expected := true

	got := cuttableWithoutMargin(materialWidth, diameter, spacing)
	if got != expected {
		t.Errorf("cuttableWithoutMargin(%d, %d, %d) == %v, want %v", materialWidth, diameter, spacing, got, expected)
	}

	materialWidth = 1000
	diameter = 100
	spacing = 60
	expected = false

	got = cuttableWithoutMargin(materialWidth, diameter, spacing)
	if got != expected {
		t.Errorf("cuttableWithoutMargin(%d, %d, %d) == %v, want %v", materialWidth, diameter, spacing, got, expected)
	}
}

func TestPossibleCutDimensions(t *testing.T) {
	cases := []struct {
		materialWidth int
		diameter      int
		expected      []CutDimensions
	}{
		{100, 40, []CutDimensions{{Diameter: 40, Spacing: 20, Pieces: 1}}},
		{900, 100, []CutDimensions{{Diameter: 100, Spacing: 60, Pieces: 5}, {Diameter: 100, Spacing: 100, Pieces: 4}}},
		{800, 200, []CutDimensions{{Diameter: 200, Spacing: 100, Pieces: 2}}},
		{1112, 12, []CutDimensions{{Diameter: 12, Spacing: 88, Pieces: 11}}},
	}

	for i, c := range cases {
		name := fmt.Sprintf("Case #%d", i)
		t.Run(name, func(t *testing.T) {
			assertCuts(t, c.materialWidth, c.diameter, c.expected)
		})
	}
}

func assertCuts(t *testing.T, materialWidth int, diameter int, expected []CutDimensions) {
	t.Helper()
	got := PossibleCutDimensions(materialWidth, diameter, 10)
	for _, e := range expected {
		if slices.Contains(got, e) == false {
			t.Errorf("it should contain %v but it didn't", e)
			t.Errorf("possibleCutDimensions(%d, %d) == %v, want %v", materialWidth, diameter, got, expected)
		}
	}
}
