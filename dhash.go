/*

Implements the dhash algorithm from http://archive.is/NFLVW

dhash is an image hashing algorithm that generates a unique
signature from an image's gradients.

The image is first grayscaled to reduce every RGB pixel set to the same value.
Then, it is resized down to 'hashLen' size, with one of the sides 1px
larger than the other (the width for horizontalGradient(), and the
height for verticalGradient()).
Finally, the gradient difference is calculated. If the current pixel is
less than the next one, a '1' is appended to the BitArray. Otherwise,
a '0' is appended.

Dhash() returns the concatenated result of horizontalGradient() & verticalGradient()
DhashHorizontal() performs only a horizontal gradient diff
DhashVertical() performs only a vertical gradient diff

TODO consider changing hashLen to reflect the length of the hash,
  instead of the length of the downscaled image
TODO test that simliar images return similar hashes
TODO Ensure Godoc is correct. BitArray is usually private
TODO Add travis.yml file
TODO then, tag as version 1.0
TODO Phash? Every new package gets a branch until testing is done

go test -v -covermode=count

*/

package imagehash

import (
  "image"
  "github.com/disintegration/imaging"
)


// Dhash returns the concatenated result of horizontalGradient() & verticalGradient()
func Dhash(img image.Image, hashLen int) ([]byte, error) {
  imgGray := imaging.Grayscale(img) // Grayscale image first for performance

  // Calculate both horizontal and vertical gradients
  horiz, err1 := horizontalGradient(imgGray, hashLen)
  vert, err2 := verticalGradient(imgGray, hashLen)

  if err1 != nil { return nil, err1 }
  if err2 != nil { return nil, err2 }

  // Return the concatenated horizontal and vertical hash
  return append(horiz, vert...), nil
}


// DhashHorizontal returns the result of a horizontal gradient diff
func DhashHorizontal(img image.Image, hashLen int) ([]byte, error) {
  imgGray := imaging.Grayscale(img) // Grayscale image first
  return horizontalGradient(imgGray, hashLen) // horizontal diff gradient
}


// DhashVertical returns the result of a vertical gradient diff
func DhashVertical(img image.Image, hashLen int) ([]byte, error) {
  imgGray := imaging.Grayscale(img) // Grayscale image first
  return verticalGradient(imgGray, hashLen) // vertical diff gradient
}


// horizontalGradient performs a horizontal gradient diff on a grayscaled image
func horizontalGradient(img image.Image, hashLen int) ([]byte, error) {
  // Width and height of the scaled-down image
  width, height := hashLen + 1, hashLen

  // Downscale the image by 'hashLen' amount for a horizonal diff.
  res := imaging.Resize(img, width, height, imaging.Lanczos)

  // Create a new bitArray
  bitArray,err := NewBitArray(hashLen * hashLen)
  if err != nil { return nil, err }

  var prev uint32 // Variable to store the previous pixel value

  // Calculate the horizonal gradient difference
  for y := 0; y < height; y++ {
    for x := 0; x < width; x++ {
      // Since the image is grayscaled, r = g = b
      r,_,_,_ := res.At(x,y).RGBA() // Get the pixel at (x,y)

      // If this is not the first value of the current row, then
      // compare the gradient difference from the previous one
      if x > 0 {
        if prev < r {
          bitArray.AppendBit(1) // if it's smaller, append '1'
        } else {
          bitArray.AppendBit(0) // else append '0'
        }
      }
      prev = r // Set this current pixel value as the previous one
    }
  }
  return bitArray.GetArray(), nil
}


// verticalGradient performs a vertical gradient diff on a grayscaled image
func verticalGradient(img image.Image, hashLen int) ([]byte, error) {
  // Width and height of the scaled-down image
  width, height := hashLen, hashLen + 1

  // Downscale the image by 'hashLen' amount for a vertical diff.
  res := imaging.Resize(img, width, height, imaging.Lanczos)

  // Create a new bitArray
  bitArray,err := NewBitArray(hashLen * hashLen)
  if err != nil { return nil, err }

  var prev uint32 // Variable to store the previous pixel value

  // Calculate the vertical gradient difference
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      // Since the image is grayscaled, r = g = b
      r,_,_,_ := res.At(x,y).RGBA() // Get the pixel at (x,y)

      // If this is not the first value of the current column, then
      // compare the gradient difference from the previous one
      if y > 0 {
        if prev < r {
          bitArray.AppendBit(1) // if it's smaller, append '1'
        } else {
          bitArray.AppendBit(0) // else append '0'
        }
      }
      prev = r // Set this current pixel value as the previous one
    }
  }
  return bitArray.GetArray(), nil
}
