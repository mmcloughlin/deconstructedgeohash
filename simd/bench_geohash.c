#include <stdio.h>
#include <inttypes.h>
#include <assert.h>

#include "geohash.h"
#include "testvector.h"
#include "benchmark.h"

void encode_int_array(size_t n, double *lat, double *lng, uint64_t *output)
{
  assert((n % BATCH_SIZE) == 0);
  for(int i = 0; i < NUM_TEST_VECTORS; i += BATCH_SIZE) {
    encode_int(lat + i, lng + i, output + i);
  }
}

int main()
{
  double lat[NUM_TEST_VECTORS];
  double lng[NUM_TEST_VECTORS];
  uint64_t hash[NUM_TEST_VECTORS];

  for(int i = 0; i < NUM_TEST_VECTORS; i++) {
    lat[i] = test_vectors[i].lat;
    lng[i] = test_vectors[i].lng;
  }

  BEST_TIME_NOCHECK(encode_int_array(NUM_TEST_VECTORS, lat, lng, hash), /* pre */, 8, NUM_TEST_VECTORS, 1);

  for(int i = 0; i < NUM_TEST_VECTORS; i++) {
    assert(hash[i] == test_vectors[i].hash);
  }

  return 0;
}

