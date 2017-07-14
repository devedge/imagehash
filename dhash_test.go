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
// be a multiple of 8, and non-zero


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
