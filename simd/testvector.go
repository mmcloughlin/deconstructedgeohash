// +build ignore

package main

import (
	"fmt"

	"github.com/mmcloughlin/deconstructedgeohash"
	"github.com/mmcloughlin/spherand"
)

const NumTestVectors = 1 << 8

func main() {
	fmt.Println("#include <stdint.h>")
	fmt.Println("typedef struct { double lat; double lng; uint64_t hash; } test_vector_t;")

	fmt.Printf("#define NUM_TEST_VECTORS (%d)\n", NumTestVectors)
	fmt.Println("test_vector_t test_vectors[NUM_TEST_VECTORS] = {")

	for i := 0; i < NumTestVectors; i++ {
		lat, lng := spherand.Geographical()
		hash := deconstructedgeohash.EncodeInt(lat, lng)
		fmt.Printf("\t{.lat = %v, .lng = %v, .hash = UINT64_C(0x%016x)},\n", lat, lng, hash)
	}

	fmt.Println("};")
}
