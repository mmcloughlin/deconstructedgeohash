package deconstructedgeohash

import (
	"math"
)

// Encode encodes the point (lat, lng) to a string geohash.
func Encode(lat, lng float64) string {
	return Base32Encode(EncodeInt(lat, lng) >> 4)
}

// EncodeInt encodes the point (lat, lng) to a 64-bit integer geohash.
func EncodeInt(lat, lng float64) uint64 {
	return Interleave(Quantize(lat, lng))
}

// EncodeIntAsm implements integer geohash in assembly.
func EncodeIntAsm(lat, lng float64) uint64

// EncodeIntSimd encodes 4 points at once.
func EncodeIntSimd(lat, lng []float64, hash []uint64)

// Quantize maps latitude and longitude to 32-bit integers.
func Quantize(lat, lng float64) (lat32 uint32, lng32 uint32) {
	lat32 = uint32(math.Ldexp((lat+90.0)/180.0, 32))
	lng32 = uint32(math.Ldexp((lng+180.0)/360.0, 32))
	return
}

// QuantizeLatAsm implements latitude quantization in assembly.
func QuantizeLatAsm(lat float64) uint32

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

// InterleaveAsm implements Interleave with the PDEP instruction.
func InterleaveAsm(x, y uint32) uint64

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

// QuantizeLatBits maps latitude to the range [1,2] and returns the bits of
// the floating point representation.
func QuantizeLatBits(lat float64) uint64 {
	return math.Float64bits(lat/180.0 + 1.5)
}
