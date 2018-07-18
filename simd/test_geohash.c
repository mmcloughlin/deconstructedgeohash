#include <stdio.h>
#include <inttypes.h>
#include <assert.h>

#include "geohash.h"
#include "testvector.h"

int main()
{
  double lat[NUM_TEST_VECTORS];
  double lng[NUM_TEST_VECTORS];
  uint64_t hash[NUM_TEST_VECTORS];

  for(int i = 0; i < NUM_TEST_VECTORS; i++) {
    lat[i] = test_vectors[i].lat;
    lng[i] = test_vectors[i].lng;
  }

  assert(NUM_TEST_VECTORS % BATCH_SIZE == 0);
  for(int i = 0; i < NUM_TEST_VECTORS; i += 4) {
    encode_int(lat + i, lng + i, hash + i);
  }

  for(int i = 0; i < NUM_TEST_VECTORS; i++) {
    if(hash[i] != test_vectors[i].hash) {
      printf("FAIL hash[%d] = %016" PRIx64 "\texpect %016" PRIx64 "\n", i, hash[i], test_vectors[i].hash);
    }
  }

  printf("pass (%d vectors)\n", NUM_TEST_VECTORS);
  return 0;
}
