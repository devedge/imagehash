/*

Average hash algorithm

*/

package imagehash

import (
  "fmt"
  "encoding/hex"
  "github.com/disintegration/imaging"
)

func AHash(img image.Image) {
  res := imaging.Grayscale(img)
  res = imaging.Resize(res, 8, 8, imaging.Lanczos)

  // As the average value is being computed, create an array
  // of pixels to optimize runtime
  var pixelArray []uint32
}
