/*

The test module for the bitarray datatype.

This tests:
1. Initializing a new valid BitArray
2. Incorrectly initializing a new BitArray with zero
3. Incorrectly initializing a new BitArray with a value under 8
4. Incorrectly initializing a new BitArray with a value above 8 but not a multiple of 8
5. Filling the BitArray with ones, using AppendBit
6. Filling the BitArray with zeros, using AppendBit
7. Failing to append an invalid bit which isn't a 0 or a 1
8. Failing to append more bits than the BitArray can contain
9. Appending fewer bits than the BitArray's full size

*/

package bitarray_test

import (
  "bytes"
  "testing"
  "../bitarray"
)


// Test initializing the bit array, which should contain only zeros
// This not only tests NewBitArray, but also the getter function, GetArray
func TestNewBitArrayInit(t *testing.T)  {
  numberBits := 32
  ba,err := bitarray.NewBitArray(numberBits)
  exp := []byte{0x00, 0x00, 0x00, 0x00}
  act := ba.GetArray()

  if bytes.Compare(exp, act) != 0 {
    t.Errorf("init test [%x] failed: [%x]", exp, act)
  } else if err != nil {
    t.Errorf("init test failed with error", err)
  }
}


// Test initializing a bit array incorrectly with 0 bits
func TestFailNewBitArrayInitZero(t *testing.T)  {
  numberBits := 0
  _,err := bitarray.NewBitArray(numberBits)

  if err == nil {
    t.Errorf("zero bits init test didn't fail")
  }
}


// Test initializing a bit array incorrectly with the number of bits
// less than 8 but not zero
func TestFailNewBitArrayLessThanEight(t *testing.T)  {
  numberBits := 5
  _,err := bitarray.NewBitArray(numberBits)

  if err == nil {
    t.Errorf("less than 8 bits init test didn't fail")
  }
}


// Test initializing a bit array incorrectly with the number of bits
// greater than 8 but not a multiple of eight
func TestFailNewBitArrayGreaterThanEight(t *testing.T)  {
  numberBits := 12
  _,err := bitarray.NewBitArray(numberBits)

  if err == nil {
    t.Errorf("greater than 8 bits init test didn't fail")
  }
}


// Test filing the bit array with ones
func TestAppendOne(t *testing.T)  {
  numberBits := 32
  ba,err := bitarray.NewBitArray(numberBits)
  exp := []byte{0xFF, 0xFF, 0xFF, 0xFF}

  // For every bit, append a 1
  for i := 0; i < numberBits; i++ {
    ba.AppendBit(1)
  }

  act := ba.GetArray()

  if bytes.Compare(exp, act) != 0 {
    t.Errorf("AppendOne test [%x] failed: [%x]", exp, act)
  } else if err != nil {
    t.Errorf("AppendOne test failed with error", err)
  }
}


// Test filing the bit array with zeros
func TestAppendZero(t *testing.T)  {
  numberBits := 32
  ba,err := bitarray.NewBitArray(numberBits)
  exp := []byte{0x00, 0x00, 0x00, 0x00}

  // For every bit, append a 0
  for i := 0; i < numberBits; i++ {
    ba.AppendBit(0)
  }

  act := ba.GetArray()

  if bytes.Compare(exp, act) != 0 {
    t.Errorf("AppendZero test [%x] failed: [%x]", exp, act)
  } else if err != nil {
    t.Errorf("AppendZero test failed with error", err)
  }
}


// Test appending an invalid bit value other than 0 or 1
func TestAppendInvalidBit(t *testing.T) {
  numberBits := 32
  ba,err := bitarray.NewBitArray(numberBits)

  // Ensure that the initialization didn't fail
  if err != nil {
    t.Errorf("AppendInvalidBit test failed init with error", err)
    return
  }

  // Append an invalid value of 24
  err = ba.AppendBit(24)

  if err == nil {
    t.Errorf("AppendInvalidBit test didn't fail")
  }
}


// Test filling the array past its capacity, which should fail
func TestOverfilledBitArray(t *testing.T)  {
  numberBits := 32
  fillamount := numberBits + 2
  ba,err := bitarray.NewBitArray(numberBits)

  // Ensure that the initialization didn't fail
  if err != nil {
    t.Errorf("overfill test failed with error", err)
    return
  }

  // Fill up only 13 of the 32 bits
  for i := 0; i < fillamount; i++ {
    err = ba.AppendBit(1)
  }

  // If overfilling the bits didn't fail, throw error
  if err == nil {
    t.Errorf("overfill test didn't fail")
  }
}


// Test filling the array below its capacity, which shouldn't fail
func TestPartlyFullBitArray(t *testing.T)  {
  numberBits := 32
  fillamount := 13
  ba,err := bitarray.NewBitArray(numberBits)
  exp := []byte{0xFF, 0xF8, 0x00, 0x00}

  // Fill up only 13 of the 32 bits
  for i := 0; i < fillamount; i++ {
    ba.AppendBit(1)
  }

  act := ba.GetArray()

  if bytes.Compare(exp, act) != 0 {
    t.Errorf("partial fill test [%x] failed: [%x]", exp, act)
  } else if err != nil {
    t.Errorf("partial fill test failed with error", err)
  }
}
