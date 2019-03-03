package q3

import (
	"testing"
)

func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}

func sliceEqual(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx, val := range s1 {
		if s2[idx] != val {
			return false
		}
	}

	return true
}

func TestGetPrimes(t *testing.T) {
	res := GetPrimes(10)
	expectedRes := []int{2, 3, 5, 7}
	if !sliceEqual(res, expectedRes) {
		t.Errorf("%v(actual) != %v(expected)", res, expectedRes)
	}
}
