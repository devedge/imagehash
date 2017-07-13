/*

This package implements a 'bit array', where bits can be added 
to a byte array one by one.
This speeds up and simplifies creating a binary signature for
an image, which can later be represented as a hex string using:
  hex.EncodeToString(imagehash)

*/

package bitarray

import (
  "errors"
  "strconv"
)


// Struct to simplify appending bits to a byte array,
// from left to right
type BitArray struct {
  byteArray []byte
  max int
  mask0 byte
  mask1 byte
  arrayIdx int
  bitIdx uint
}


/**
 * Constructor for the BitArray struct. It must be initialized
 * with a non-zero int that is a multiple of 8.
 * Usage:
 *    bitArray,err := bitops.NewBitArray(64)
 *
 * @method NewBitArray
 * @param  {int}        numBits The number of bits, as an int
 * @return {BitArray}   The AppendBit struct
 * @return {error}      If there is an error
 */
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


/**
 * Append a bit to the byte array from left to right
 * @method appendBit
 * @param  {int} bit  Append a one with '1', and zero with '0'
 * @return error      An error, if any occured
 */
func (ab *BitArray) AppendBit(bit int) error {
  if ab.arrayIdx == ab.max {
    return errors.New("cannot contine to append to a full byte array")
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
    // bit to be set will be on the right.
    ab.bitIdx--
  } else {
    // The last bit in the current byte has been set, so increment
    // the index into the byte array, and reset the bit index.
    ab.arrayIdx++
    ab.bitIdx = 7
  }

  return nil
}


/**
 * Returns the current byte array
 * @method getArray
 * @return {[]byte]}  The current byte array
 */
func (ab BitArray) GetArray() []byte {
  return ab.byteArray
}
