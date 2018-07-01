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

func TestQuantizeLatAsm(t *testing.T) {
	for trial := 0; trial < NumTrials(); trial++ {
		lat := RandomLat()
		expect, _ := Quantize(lat, 0)
		got := QuantizeLatAsm(lat)
		if math.Abs(float64(got-expect)) > 1 {
			t.Errorf("got=%08x expect=%08x delta=%d", got, expect, got-expect)
		}
	}
}

func TestInterleaveAsm(t *testing.T) {
	for trial := 0; trial < NumTrials(); trial++ {
		x, y := rand.Uint32(), rand.Uint32()
		expect := Interleave(x, y)
		got := InterleaveAsm(x, y)
		if expect != got {
			t.Errorf("got=%016x expect=%016x", got, expect)
		}
	}
}
