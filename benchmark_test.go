package deconstructedgeohash

import "testing"

var lat, lng = 40.463833, -79.972422

func BenchmarkEncodeInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncodeInt(lat, lng)
	}
}
