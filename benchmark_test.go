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
