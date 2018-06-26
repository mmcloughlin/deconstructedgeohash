package deconstructedgeohash

import (
	"math"
)

func Quantize(lat, lng float64) (lat32 uint32, lng32 uint32) {
	lat32 = uint32(math.Ldexp((lat+90.0)/180.0, 32))
	lng32 = uint32(math.Ldexp((lng+180.0)/360.0, 32))
	return
}

// Spread out the 32 bits of x into 64 bits, where the bits of x occupy even
// bit positions.
func Spread(x uint32) uint64 {
	X := uint64(x)
	X = (X | (X << 16)) & 0x0000ffff0000ffff
	X = (X | (X << 8)) & 0x00ff00ff00ff00ff
	X = (X | (X << 4)) & 0x0f0f0f0f0f0f0f0f
	X = (X | (X << 2)) & 0x3333333333333333
	X = (X | (X << 1)) & 0x5555555555555555
	return X
}

// Interleave the bits of x and y. In the result, x and y occupy even and odd
// bitlevels, respectively.
func Interleave(x, y uint32) uint64 {
	return Spread(x) | (Spread(y) << 1)
}

// Base32Encode the bits of x according to the geohash alphabet.
func Base32Encode(x uint64) string {
	alpha := "0123456789bcdefghjkmnpqrstuvwxyz"
	b := [12]byte{}
	for i := 0; i < 12; i++ {
		b[11-i] = alpha[x&0x1f]
		x >>= 5
	}
	return string(b[:])
}
