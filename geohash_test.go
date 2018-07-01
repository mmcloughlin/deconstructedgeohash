package deconstructedgeohash

import (
	"math"
	"math/rand"
	"testing"
)

func NumTrials() int {
	if testing.Short() {
		return 1 << 5
	}
	return 1 << 25
}

func RandomLat() float64 {
	return -90 + 180*rand.Float64()
}

func TestEncode(t *testing.T) {
	if Encode(-25.382708, -49.265506) != "6gkzwgjzn820" {
		t.Fail()
	}
}

func TestQuantize(t *testing.T) {
	lat32, lng32 := Quantize(27.988056, 86.925278)
	if lat32 != 0xa7ce23e4 || lng32 != 0xbdd04391 {
		t.Fail()
	}
}

func TestQuantizeLat(t *testing.T) {
	for trial := 0; trial < NumTrials(); trial++ {
		lat := RandomLat()
		expect, _ := Quantize(lat, 0)
		lat32 := QuantizeLat(lat)
		if math.Abs(float64(lat32-expect)) > 1 {
			t.Errorf("lat32=%08x expect=%08x delta=%d", lat32, expect, lat32-expect)
		}
	}
}
