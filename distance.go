package imagehash

// GetDistance returns the hamming distance between two hashs
func GetDistance(hash1, hash2 []byte) int {
	distance := 0

	// Difines the shorter and the longer hash
	var shorter int
	var longer int
	if len(hash1) <= len(hash2) {
		shorter = len(hash1)
		longer = len(hash2)
	} else {
		shorter = len(hash2)
		longer = len(hash1)
	}

	// Check the distance
	for i := 0; i < shorter; i++ {
		if hash1[i] != hash2[i] {
			distance++
		}
	}

	// Add the deferance not computable because the two elements
	// don't have the same length.
	distance += longer - shorter

	return distance
}

// GetDistanceMaxRange returns the distance between two hashs
func GetDistanceMaxRange(hash1, hash2 []byte) int {
	if len(hash1) >= len(hash2) {
		return len(hash1)
	}
	return len(hash2)
}
