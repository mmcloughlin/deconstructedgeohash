// +build ignore

package main

import (
	"fmt"

	"github.com/mmcloughlin/geohash"
)

const NumTestVectors = 4

func main() {
	fmt.Println("#include <stdint.h>")
	fmt.Println("typedef struct { double lat; double lng; uint64_t hash; } test_vector_t;")

	fmt.Printf("#define NUM_TEST_VECTORS (%d)\n", NumTestVectors)
	fmt.Println("test_vector_t test_vectors[NUM_TEST_VECTORS] = {")

	for i := uint64(1); i <= 4; i++ {
		hash := i * uint64(0x1111111111111111)
		lat, lng := geohash.BoundingBoxInt(hash).Center()
		fmt.Printf("\t{.lat = %f, .lng = %f, .hash = UINT64_C(0x%016x)},\n", lat, lng, hash)
	}

	fmt.Println("};")
}
