package deconstructedgeohash

import "testing"

// BenchmarkEncodeInt benchmarks integer geohash encoding.
func BenchmarkEncodeInt(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeInt(lat, lng)
	}
}

// BenchmarkEncodeIntAsm benchmarks integer geohash encoding.
func BenchmarkEncodeIntAsm(b *testing.B) {
	var lat, lng = 40.463833, -79.972422
	for i := 0; i < b.N; i++ {
		EncodeIntAsm(lat, lng)
	}
}
