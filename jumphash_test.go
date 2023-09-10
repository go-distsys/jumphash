package jumphash

import (
	"testing"
)

// See: https://github.com/google/guava/blob/7d6d372f2428e2f8d07ee3162d10e81859e0c83c/android/guava-tests/test/com/google/common/hash/HashingTest.java
func TestHash(t *testing.T) {
	golden100 := []int{0, 55, 62, 8, 45, 59, 86, 97, 82, 59, 73, 37, 17, 56, 86, 21, 90, 37, 38, 83}

	for i, want := range golden100 {
		res := Hash(uint64(i), 100)
		mustEqual(t, res, want)
	}

	testCases := []struct {
		key     uint64
		buckets int
		res     int
	}{
		{key: 10863919174838991, buckets: 11, res: 6},
		{key: 2016238256797177309, buckets: 11, res: 3},
		{key: 1673758223894951030, buckets: 11, res: 5},
		{key: 2, buckets: 100001, res: 80343},
		{key: 2201, buckets: 100001, res: 22152},
		{key: 2202, buckets: 100001, res: 15018},
		{key: 42, buckets: -1, res: 0},
	}

	for _, tt := range testCases {
		res := Hash(tt.key, tt.buckets)
		mustEqual(t, res, tt.res)
	}
}

func mustEqual(tb testing.TB, have, want int) {
	tb.Helper()

	if have != want {
		tb.Errorf("\nhave: %v\nwant: %v", have, want)
	}
}
