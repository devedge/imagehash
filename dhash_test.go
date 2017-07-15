/*

Testing suite for the dhash algorithm.

1. Test the white dhash
2. Test the white horizontal dhash
3. Test the white vertical dhash
4. Test an invalid dhash length
5. Test an invalid horizontal dhash length
6. Test an invalid vertical dhash length
7. Test that the dhash of lena_512 matches the precomputed one
8. Test that the dhash of a 512px image and 256px image are identical

*/

package imagehash

import (
  "bytes"
  "testing"
)


// Test computing the Dhash of a white image. The
// resultant byte array should be all zeros
func TestWhiteDhash(t *testing.T)  {
  src,_ := OpenImg("./testdata/white_512.png")
  hash,err := Dhash(src, 8)
  exp := make([]byte, 16) // initialize a byte array full of zeros

  if bytes.Compare(exp, hash) != 0 {
    t.Errorf("white dhash test [%x] failed: [%x]", exp, hash)
  } else if err != nil {
    t.Errorf("white dhash test failed with error:", err)
  }
}

// Test the horizontal dhash wrapper function using a white image. The
// resultant byte array should be all zeros
func TestWhiteDhashHorizontal(t *testing.T)  {
  src,_ := OpenImg("./testdata/white_512.png")
  hash,err := DhashHorizontal(src, 8)
  exp := make([]byte, 8) // initialize a byte array full of zeros

  if bytes.Compare(exp, hash) != 0 {
    t.Errorf("white horizontal dhash test [%x] failed: [%x]", exp, hash)
  } else if err != nil {
    t.Errorf("white horizontal dhash test failed with error:", err)
  }
}

// Test the vertical dhash wrapper function using a white image. The
// resultant byte array should be all zeros
func TestWhiteDhashVertical(t *testing.T)  {
  src,_ := OpenImg("./testdata/white_512.png")
  hash,err := DhashVertical(src, 8)
  exp := make([]byte, 8) // initialize a byte array full of zeros

  if bytes.Compare(exp, hash) != 0 {
    t.Errorf("white vertical dhash test [%x] failed: [%x]", exp, hash)
  } else if err != nil {
    t.Errorf("white vertical dhash test failed with error:", err)
  }
}


// Test passing in an invalid hashLen. The result of hashLen * hashLen must
// be a multiple of 8, and non-zero.
// (These three following tests only check 0, since bitarray_test.go cover
// all other cases)
func TestDhashZeroHashLen(t *testing.T) {
  src,_ := OpenImg("./testdata/white_512.png")
  _,err := Dhash(src, 0)

  if err == nil {
    t.Errorf("zero dhash hashLen didn't fail")
  }
}
func TestDhashHorizontalZeroHashLen(t *testing.T) {
  src,_ := OpenImg("./testdata/white_512.png")
  _,err := DhashHorizontal(src, 0)

  if err == nil {
    t.Errorf("zero horizontal dhash hashLen didn't fail")
  }
}
func TestDhashVerticalZeroHashLen(t *testing.T) {
  src,_ := OpenImg("./testdata/white_512.png")
  _,err := DhashVertical(src, 0)

  if err == nil {
    t.Errorf("zero vertical dhash hashLen didn't fail")
  }
}


// Test that the lena_512.png image returns the expected result.
// The value of 'exp' has already been manually computed for the
// horizontal gradient.
func TestLenaHash(t *testing.T) {
  src,_ := OpenImg("./testdata/lena_512.png")
  hash,err := DhashHorizontal(src, 8)
  exp := []byte{0x76, 0x70, 0x79, 0x5b, 0x33, 0x13, 0x5a, 0x38}

  if bytes.Compare(exp, hash) != 0 {
    t.Errorf("basic lena_512 dhash test [%x] failed: [%x]", exp, hash)
  } else if err != nil {
    t.Errorf("basic lena_512 dhash test failed with error:", err)
  }
}


// Test that lena_512.png and lena_256.png return identical dhash results.
func TestSimilarLenaDash(t *testing.T)  {
  lena512,_ := OpenImg("./testdata/lena_512.png")
  lena256,_ := OpenImg("./testdata/lena_256.png")
  hashlena512,err1 := Dhash(lena512, 8)
  hashlena256,err2 := Dhash(lena256, 8)

  if bytes.Compare(hashlena512, hashlena256) != 0 {
    t.Errorf("similar lena dhash test [%x] failed: [%x]", hashlena512, hashlena256)
  } else if err1 != nil {
    t.Errorf("similar lena dhash test failed with error:", err1)
  } else if err2 != nil {
    t.Errorf("similar lena dhash test failed with error:", err2)
  }
}
