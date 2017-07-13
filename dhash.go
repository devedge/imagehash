package imagehash

import (
  "image"
  "./bitarray"
  "github.com/disintegration/imaging"
)


func Dhash(img image.Image, hashLen int) ([]byte, error) {
  imgGray := imaging.Grayscale(img) // Grayscale image first for performance

  // Calculate both horizontal and vertical gradients
  horiz, err1 := horizontalGradient(imgGray, hashLen)
  vert, err2 := verticalGradient(imgGray, hashLen)

  if err1 != nil {
    return nil, err1
  } else if err2 != nil {
    return nil, err2
  }

  // Return the horizontal hash with the vertical one appended
  return append(horiz, vert...), nil
}


func DhashHorizontal(img image.Image, hashLen int) ([]byte, error) {
  imgGray := imaging.Grayscale(img) // Grayscale image first
  horiz, err := horizontalGradient(imgGray, hashLen) // horizontal diff gradient
  return horiz, err
}


func DhashVertical(img image.Image, hashLen int) ([]byte, error) {
  imgGray := imaging.Grayscale(img) // Grayscale image first
  horiz, err := verticalGradient(imgGray, hashLen) // horizontal diff gradient
  return horiz, err
}


func horizontalGradient(img image.Image, hashLen int) ([]byte, error) {
  // Width and height of the scaled-down image
  width, height := hashLen + 1, hashLen

  // Downscale the image by 'hashLen' amount for a horizonal diff.
  res := imaging.Resize(img, width, height, imaging.Lanczos)

  // Create a new bitArray
  bitArray,err := bitarray.NewBitArray(hashLen * hashLen)
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


func verticalGradient(img image.Image, hashLen int) ([]byte, error) {
  // Width and height of the scaled-down image
  width, height := hashLen, hashLen + 1

  // Downscale the image by 'hashLen' amount for a horizonal diff.
  res := imaging.Resize(img, width, height, imaging.Lanczos)

  // Create a new bitArray
  bitArray,err := bitarray.NewBitArray(hashLen * hashLen)
  if err != nil { return nil, err }

  var prev uint32 // Variable to store the previous pixel value

  // Calculate the horizonal gradient difference
  for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
      // Since the image is grayscaled, r = g = b
      r,_,_,_ := res.At(x,y).RGBA() // Get the pixel at (x,y)

      // If this is not the first value of the current row, then
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
