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
	return 1 << 26
}

func RandomLat() float64 {
	return -90 + 180*rand.Float64()
}

func RandomLng() float64 {
	return -180 + 360*rand.Float64()
}

func TestEncode(t *testing.T) {
	if Encode(-25.382708, -49.265506) != "6gkzwgjzn820" {
		t.Fail()
	}
}

func TestEncodeInt(t *testing.T) {
	type IntEncoder func(float64, float64) uint64
	for _, encoder := range []IntEncoder{EncodeInt, EncodeIntAsm} {
		if encoder(-25.382708, -49.265506) != 0x33e5fe3e3fa2040f {
			t.FailNow()
		}
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

func TestEncodeIntSimd(t *testing.T) {
	lat := make([]float64, 4)
	lng := make([]float64, 4)
	hash := make([]uint64, 4)
	for trial := 0; trial < NumTrials(); trial += 4 {
		for i := 0; i < 4; i++ {
			lat[i], lng[i] = RandomLat(), RandomLng()
		}

		EncodeIntSimd(lat, lng, hash)

		for i := 0; i < 4; i++ {
			expect := EncodeIntAsm(lat[i], lng[i])
			if expect != hash[i] {
				t.Fatalf("lat=%f\tlng=%f\tgot=%016x\texpect=%016x\n", lat[i], lng[i], hash[i], expect)
			}
		}
	}
}
