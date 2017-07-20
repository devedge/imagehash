# imagehash
[![Build Status](https://travis-ci.org/devedge/imagehash.svg?branch=master)](https://travis-ci.org/devedge/imagehash)
[![GoDoc](https://godoc.org/github.com/devedge/imagehash?status.svg)](https://godoc.org/github.com/devedge/imagehash)
[![Coverage](https://img.shields.io/badge/coverage-98.7-brightgreen.svg)](https://gocover.io/github.com/devedge/imagehash)

Golang implementation of image hashing algorithms.


## Install:

`go get -u github.com/devedge/imagehash`


## Usage

There are currently two image hashing algorithms implemented:
 - [dhash](#dhash) - difference/gradient hash
 - [ahash](#ahash) - average hash

To hash an image, it must be opened using `OpenImg`, a wrapper around `imaging`'s image decoding function.
```go
src,err := imagehash.OpenImg("./testdata/lena_512.png")
```

 - more general usage information can be found [in the example section](#examples)


## dhash

This is an implementation of the `dhash` algorithm [found here](http://archive.is/NFLVW) (archived link), and also [implemented in Python  here](https://github.com/JohannesBuchner/imagehash).

It generates a unique signature for an image based on the gradient difference between pixels.


```go
// The hashes are returned as byte arrays
//
// Calculate both horizontal & vertical gradients, then return them
// concatenated together as <horizontal><vertical>
hash,err  := imagehash.Dhash(src, hashLen)

// Calculate only the horizontal gradient difference
hashH,err := imagehash.DhashHorizontal(src, hashLen)

// Calculate only the vertical gradient difference
hashV,err := imagehash.DhashVertical(src, hashLen)
```

#### Using dhash:

```go
package main

import (
  "fmt"
  "encoding/hex"  // To transform the byte array to hex
  "github.com/devedge/imagehash"
)

func main() {
  src,_ := imagehash.OpenImg("./testdata/lena_512.png")

  // The length of a downscaled side. It must be > 8, and
  // (hashLen * hashLen) must be a multiple of 8
  hashLen := 8
  // A value of 8 will return 64 bits, or 8 bytes / 16 hex characters
  // (64 bits = 8 bits length * 8 bits width)

  hash,_ := imagehash.Dhash(src, hashLen)
  hashH,_ := imagehash.DhashHorizontal(src, hashLen)
  hashV,_ := imagehash.DhashVertical(src, hashLen)

  fmt.Println("dhash:           ", hex.EncodeToString(hash))
  fmt.Println("Horizontal dhash:", hex.EncodeToString(hashH))
  fmt.Println("Vertical dhash:  ", hex.EncodeToString(hashV))
}
```

#### Implementation:

First, the image is grayscaled:
<br>

![grayscale](doc/grayscale.png)

To calculate the horizontal gradient difference, the image is resized down, using the `hashLen` variable.

In this example, `hashLen = 8`, so the image is scaled down to `9x8px`. Then, if `pixel[x] < pixel[x+1]`, a `1` is appended to a byte array; otherwise, a `0`. This results in 8 bits of data per row, for 8 columns, or 64 bits total:
<br>

![dhashprocess](doc/process.png)

This array of 1s and 0s is then flattened, and returned as a byte array: <br>
`0111011001110000011110010101101100110011000100110101101000111000`

Which can also be represented in hex as `7670795b33135a38` using `hex.EncodeToString(result)`
<br>

Conversely, to obtain a vertical diff, the image would be scaled down to `8x9px`, and the diff matrix would be the result of `pixel[y] < pixel[y+1]`.


## ahash

This algorithm returns a hash based on the average pixel value.

As with dhash, it also grayscales and resizes the image down, using the 'hashLen' value. Then, it finds the average value of the resultant pixels. Finally, it iterates over the pixels, and if one is greater than the average, a `1` is appended to the returned result; a `0` otherwise.


```go
// The hash is returned as a byte array
hash,err := imagehash.Ahash(src, hashLen)
```


## Examples

The Hamming distance between two byte arrays can be determined using a package like [hamming](https://github.com/steakknife/hamming):

```go
package main

import (
  "fmt"
  "encoding/hex"
  "github.com/devedge/imagehash"
  "github.com/steakknife/hamming"
)

func main() {
  src512,_ := imagehash.OpenImg("./testdata/lena_512.png")
  src256,_ := imagehash.OpenImg("./testdata/lena_256.png")
  srcInv,_ := imagehash.OpenImg("./testdata/lena_inverted_512.png")

  hash512,_ := imagehash.Dhash(src512, 8)
  hash256,_ := imagehash.Dhash(src256, 8)
  hashInv,_ := imagehash.Dhash(srcInv, 8)

  // Hamming distance of 0, since the images are simply different sizes
  fmt.Println("'lena_512.png' dhash:", hex.EncodeToString(hash512))
  fmt.Println("'lena_256.png' dhash:", hex.EncodeToString(hash256))
  fmt.Println("The Hamming distance between these:", hamming.Bytes(hash512, hash256))
  fmt.Println()

  // Completely different dhash, since an inverted image has a completely
  // different gradient colorscheme
  fmt.Println("'lena_512.png' dhash:         ", hex.EncodeToString(hash512))
  fmt.Println("'lena_inverted_512.png' dhash:", hex.EncodeToString(hashInv))
  fmt.Println("The Hamming distance between these:", hamming.Bytes(hash512, hashInv))
}
```


## Dependencies:
* [imaging](https://github.com/disintegration/imaging) - Simple Go image processing package
