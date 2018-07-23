package deconstructedgeohash

import "testing"

// BenchmarkEncodeInt benchmarks integer geohash encoding.
func BenchmarkEncodeInt(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeInt(lat, lng)
	}
}

// BenchmarkEncodeIntAsm benchmarks assembly integer geohash encoding.
func BenchmarkEncodeIntAsm(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeIntAsm(lat, lng)
	}
}

// NoopAsm does nothing in assembly, to assess function call overhead.
func NoopAsm(lat, lng float64) uint64

// BenchmarkNoopAsm benchmarks the noop assembly function.
func BenchmarkNoopAsm(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		NoopAsm(lat, lng)
	}
}

// BenchmarkEncodeIntSimd benchmarks SIMD integer geohash encoding.
func BenchmarkEncodeIntSimd(b *testing.B) {
	const lat, lng = 40.463833, -79.972422
	lat4 := []float64{lat, lat, lat, lat}
	lng4 := []float64{lng, lng, lng, lng}
	hash4 := make([]uint64, 4)
	for i := 0; i < b.N; i++ {
		EncodeIntSimd(lat4, lng4, hash)
	}
}
