/*

This is an internal module to simplify adding bits one at a
time to a byte array. It is used by the dhash algorithm, where
the byte array can later be transformed into its hex representation.

Example usage:
  bitArray,err := NewBitArray(32)
  bitArray.AppendBit(1)
  fmt.Println(bitArray.GetArray())

*/

package imagehash

import (
  "errors"
  "strconv"
)


// BitArray is an internal struct used by dhash to simplify appending bits to
// a byte array from left to right. 
type BitArray struct {
  byteArray []byte
  max int
  mask0 byte
  mask1 byte
  arrayIdx int
  bitIdx uint
}


// NewBitArray is a constructor function for the BitArray struct.
// The input, 'numBits' is the number of bits this byte array will
// hold, so it must be a non-zero multiple of 8.
func NewBitArray(numBits int) (*BitArray, error) {
  // If numBits is invalid
  if (numBits == 0) || (numBits % 8 != 0) {
    return nil, errors.New("'numBits' must be a non-zero multiple of 8")
  }

  return &BitArray{
    byteArray: make([]byte, numBits / 8),
    max: numBits / 8,
    mask0: 0x00,
    mask1: 0x01,
    arrayIdx: 0,
    bitIdx: 7,
  }, nil
}


// AppendBit appends a 1 or a 0 to the byte array in the BitArray struct.
// Valid input is an int of '1' or '0', and this function cannot be called
// after the byte array has filled up.
func (ab *BitArray) AppendBit(bit int) error {
  if ab.arrayIdx == ab.max {
    return errors.New("cannot continue to append to a full byte array")
  }

  // Shift the 'mask' (bit of 1 or 0) by the proper amount to
  // fill the byte up from left to right.
  switch bit {
    case 0:
      ab.byteArray[ab.arrayIdx] |= ab.mask0 << ab.bitIdx
    case 1:
      ab.byteArray[ab.arrayIdx] |= ab.mask1 << ab.bitIdx
    default:
      return errors.New("can only append with 1 or 0, but received: " + strconv.Itoa(bit))
  }

  if ab.bitIdx > 0 {
    // Decrement the index into the current byte so the next
    // bit to set will be on the right.
    ab.bitIdx--
  } else {
    // The last bit in the current byte has been set, so increment
    // the index into the byte array, and reset the bit index.
    ab.arrayIdx++
    ab.bitIdx = 7
  }

  return nil
}


// GetArray returns the byte array in its current state. It
// can be called at any time.
func (ab BitArray) GetArray() []byte {
  return ab.byteArray
}
