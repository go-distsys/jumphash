package jumphash

// Hash consistently chooses a bucket in the range [0, numBuckets) for the given key.
// Zero or negative numBuckets are treated as numBuckets = 1.
// See: https://arxiv.org/pdf/1406.2294.pdf
func Hash(key uint64, numBuckets int) int {
	if numBuckets <= 0 {
		numBuckets = 1
	}

	var b int64 = -1
	for j := int64(0); j < int64(numBuckets); {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(int64(1)<<31) / float64((key>>33)+1)))
	}
	return int(b)
}
