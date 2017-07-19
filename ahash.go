/*

Implements the average hash algorithm from <archive link>

This algorithm returns a hash based on the average pixel value.

First, it grayscales and resizes the image down, using the 'hashLen'
value. Then, it finds the average pixel value from this image.
Finally, it iterates over the pixels, and if one is greater than the average,
a 1 is appended to the returned result; a 0 otherwise.

*/

package imagehash

import (
  "image"
  "github.com/disintegration/imaging"
)


// Ahash calculates the average hash of an image. The image is first grayscaled,
// then scaled down to "hashLen" for the width and height. Then, the average value
// of the pixels is computed, and if a pixel is above the average, a 1 is appended
// to the byte array; a 0 otherwise.
func Ahash(img image.Image, hashLen int) ([]byte, error) {
  var sum uint32                        // Sum of the pixels
  numbits := hashLen * hashLen          // Perform the hashLen^2 operation once
  bitArray,err := NewBitArray(numbits)  // Resultant byte array init
  if err != nil { return nil, err }

  // As the average is being computed, create & populate an array
  // of pixels to optimize runtime
  var pixelArray []uint32

  // Grayscale and resize
  res := imaging.Grayscale(img)
  res = imaging.Resize(res, hashLen, hashLen, imaging.Lanczos)

  // Iterate over every pixel to generate the sum.
  // Additionally, store every pixel into an array for faster re-computation
  for x := 0; x < hashLen; x++ {
    for y := 0; y < hashLen; y++ {
      r,_,_,_ := res.At(x,y).RGBA()       // r = g = b since the image is grayscaled
      sum += r                            // increment the sum
      pixelArray = append(pixelArray, r)  // append the pixel
    }
  }

  // Compute the average
  avg := sum / uint32(numbits)

  // For every pixel, check if it's below or above the average
  for _,pix := range pixelArray {
    if pix > avg {
      bitArray.AppendBit(1) // If above, append 1
    } else {
      bitArray.AppendBit(0) // else append 0
    }
  }

  return bitArray.GetArray(), nil
}
