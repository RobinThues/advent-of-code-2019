package helpers

import "testing"

func TestStringSliceToIntSlice(t *testing.T) {
	s := []string{
		"123",
	}
	i, err := StringSliceToIntSlice(s)
	if i[0] != 123 {
		t.Errorf("StringSliceToIntSlice({'123'}) = %v, wanted [123], nil", i)
	}
	if err != nil {
		t.Errorf("StringSliceToIntSlice({'123'}) error is not nil")
	}
}
