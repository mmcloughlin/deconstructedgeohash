package deconstructedgeohash

import (
	"math"
	"math/rand"
	"testing"

	"github.com/mmcloughlin/geohash"
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

func TestEncodeInt4(t *testing.T) {
	var lat, lng [4]float64
	for i := 0; i < 4; i++ {
		hash := uint64(0x1111111111111111) * uint64(i)
		box := geohash.BoundingBoxInt(hash)
		lat[i], lng[i] = box.Center()
	}

	hash := EncodeInt4(lat, lng)

	for i := 0; i < 4; i++ {
		if EncodeInt(lat[i], lng[i]) != hash[i] {
			t.Fatalf("lat=%f\tlng=%f\tgot=%016x\texpect=%016x\n", lat[i], lng[i], hash[i], EncodeInt(lat[i], lng[i]))
		}
	}
}
