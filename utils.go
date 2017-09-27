/*

Wrapper function around 'imaging''s Open() fuction, so
images can be opened directly through the 'imagehash' package
instead of requiring 2 packages for usage.

Usage:
  img,err := imagehash.OpenImg("image.jpg")

*/

package imagehash

import (
	"github.com/disintegration/imaging"
	"image"
)

// OpenImg is a wrapper aroung the Open function from 'imaging'.
// Open opens & encodes an image from the filesystem, which dhash is
// based upon.
func OpenImg(fp string) (image.Image, error) {
	return imaging.Open(fp)
}
