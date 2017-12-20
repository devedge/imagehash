/*

Testing suite for the distance definition with all implemented algorithms.
It starts with Ahash and finish with Dhash.

1. Test that distance between a 512px image and 256px image is near to zero
2. Test that distance between a image and inverted image is not too close from zero but not to far either
3. Test that distance between a image and white image is not close from zero
4. Test maximum distance between a images

*/

package imagehash

import (
	"testing"
)

// Test the distance between two size with Ahash.
func TestDistanceAHashSizes(t *testing.T) {
	src1, _ := OpenImg("./testdata/lena_512.png")
	src2, _ := OpenImg("./testdata/lena_256.png")
	hash1, _ := Ahash(src1, 8)
	hash2, _ := Ahash(src2, 8)

	dist := GetDistance(hash1, hash2)

	// The value should be 0
	if dist > 3 {
		t.Errorf("the distance between the two images is to important: %d", dist)
	}
}

// Test the distance between two size with Dhash.
func TestDistanceDHashSizes(t *testing.T) {
	src1, _ := OpenImg("./testdata/lena_512.png")
	src2, _ := OpenImg("./testdata/lena_256.png")
	hash1, _ := Dhash(src1, 8)
	hash2, _ := Dhash(src2, 8)

	dist := GetDistance(hash1, hash2)

	// The value should be 0
	if dist > 3 {
		t.Errorf("the distance between the two images is to important: %d", dist)
	}
}

// Test the distance between regular and inverted image with Ahash.
func TestDistanceAHashInverted(t *testing.T) {
	src1, _ := OpenImg("./testdata/lena_512.png")
	src2, _ := OpenImg("./testdata/lena_256.png")
	hash1, _ := Ahash(src1, 8)
	hash2, _ := Ahash(src2, 8)

	dist := GetDistance(hash1, hash2)

	// The value should be 0
	if dist > 3 {
		t.Errorf("the distance between the two images is to important: %d", dist)
	}
}

// Test the distance between regular and inverted image with Dhash.
func TestDistanceDHashInverted(t *testing.T) {
	src1, _ := OpenImg("./testdata/lena_512.png")
	src2, _ := OpenImg("./testdata/lena_inverted_512.png")
	hash1, _ := Dhash(src1, 8)
	hash2, _ := Dhash(src2, 8)

	dist := GetDistance(hash1, hash2)

	// The value should be 16
	if dist < 10 && dist > 20 {
		t.Errorf("the distance between the two images is to important: %d", dist)
	}
}

// Test the distance between regular and inverted image with Ahash.
func TestDistanceAHashWhite(t *testing.T) {
	src1, _ := OpenImg("./testdata/lena_512.png")
	src2, _ := OpenImg("./testdata/white_512.png")
	hash1, _ := Ahash(src1, 8)
	hash2, _ := Ahash(src2, 8)

	dist := GetDistance(hash1, hash2)

	// The value should be 7
	if dist < 5 {
		t.Errorf("the distance between the two images is to important: %d", dist)
	}
}

// Test the distance between regular and inverted image with Dhash.
func TestDistanceDHashWhite(t *testing.T) {
	src1, _ := OpenImg("./testdata/lena_512.png")
	src2, _ := OpenImg("./testdata/white_512.png")
	hash1, _ := Dhash(src1, 8)
	hash2, _ := Dhash(src2, 8)

	dist := GetDistance(hash1, hash2)

	// The value should be 16
	if dist < 10 {
		t.Errorf("the distance between the two images is to important: %d", dist)
	}
}

// Test the maximum distance between hashes of 8 and 16.
func TestMaximuimDistance(t *testing.T) {
	src1, _ := OpenImg("./testdata/white_512.png")
	src2, _ := OpenImg("./testdata/rand_512.png")

	hash1, _ := Ahash(src1, 8)
	hash2, _ := Ahash(src2, 8)

	dist := GetDistance(hash1, hash2)
	// The value should be 16
	if distMax := GetDistanceMaxRange(hash1, hash2); distMax != dist {
		t.Errorf("the maximum distance is not good. We have %d and it should be %d", distMax, dist)
	}

	hash1, _ = Ahash(src1, 16)
	hash2, _ = Ahash(src2, 16)

	dist = GetDistance(hash1, hash2)
	// The value should be 32
	if distMax := GetDistanceMaxRange(hash1, hash2); distMax != dist {
		t.Errorf("the maximum distance is not good. We have %d and it should be %d", distMax, dist)
	}

	hash1, _ = Ahash(src1, 32)
	hash2, _ = Ahash(src2, 16)

	dist = GetDistance(hash1, hash2)
	// The value should be 32
	if distMax := GetDistanceMaxRange(hash1, hash2); distMax != dist {
		t.Errorf("the maximum distance is not good. We have %d and it should be %d", distMax, dist)
	}
}
