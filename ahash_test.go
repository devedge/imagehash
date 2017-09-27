/*

Testing suite for the ahash algorithm.

1. Test the white ahash
2. Test an invalid ahash length
3. Test that the ahash of lena_512 matches the precomputed one
4. Test that the ahash of a 512px image and 256px image are similar

*/

package imagehash

import (
	"bytes"
	"testing"
)

// Test computing the Ahash of a white image. The resultant
// byte array should be all zeros
func TestWhiteAhash(t *testing.T) {
	src, _ := OpenImg("./testdata/white_512.png")
	hash, err := Ahash(src, 8)
	exp := make([]byte, 8) // initialize a byte array full of zeros

	if bytes.Compare(exp, hash) != 0 {
		t.Errorf("white ahash test [%x] failed: [%x]", exp, hash)
	} else if err != nil {
		t.Errorf("white ahash test failed with error:", err)
	}
}

// Test an invalid hashLen of zero. This just ensures that errors
// from BitArray get properly passed up.
func TestAhashZeroHashLen(t *testing.T) {
	src, _ := OpenImg("./testdata/white_512.png")
	_, err := Ahash(src, 0)

	if err == nil {
		t.Errorf("zero ahash hashLen didn't fail")
	}
}

// Test the Lena 512 image
func TestLenaAhash(t *testing.T) {
	src, _ := OpenImg("./testdata/lena_512.png")
	hash, err := Ahash(src, 8)
	exp := []byte{0xf3, 0x00, 0xa0, 0xe0, 0x7f, 0xe3, 0x8e, 0x3e}

	if bytes.Compare(exp, hash) != 0 {
		t.Errorf("lena_512 ahash test [%x] failed: [%x]", exp, hash)
	} else if err != nil {
		t.Errorf("lena_512 ahash test failed with error:", err)
	}
}

// Test that the lena_512 and lena_256 images return identical average hashes
func TestSimilarLenaAhash(t *testing.T) {
	lena512, _ := OpenImg("./testdata/lena_512.png")
	lena256, _ := OpenImg("./testdata/lena_256.png")
	hashlena512, err1 := Ahash(lena512, 12) // more bits to get better accuracy
	hashlena256, err2 := Ahash(lena256, 12) // more bits to get better accuracy

	if bytes.Compare(hashlena512, hashlena256) != 0 {
		t.Errorf("similar lena ahash test [%x] failed: [%x]", hashlena512, hashlena256)
	} else if err1 != nil {
		t.Errorf("similar lena ahash test failed with error:", err1)
	} else if err2 != nil {
		t.Errorf("similar lena ahash test failed with error:", err2)
	}
}
