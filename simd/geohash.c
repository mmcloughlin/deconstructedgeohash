// SIMD implementation of integer geohash encoding.
//
// Based on work by Daniel Lemire and Geoff Langdale.
// https://lemire.me/blog/2018/01/09/how-fast-can-you-bit-interleave-32-bit-integers-simd-edition/

#include <stdio.h>
#include <inttypes.h>
#include <stdlib.h>
#include <immintrin.h>
#include <assert.h>

#include "testvector.h"

// print_double prints the 4 doubles in x.
void print_double(__m256d x)
{
  double mem[4];
  _mm256_storeu_pd(mem, x);
  for(int i = 0; i < 4; i++) {
    printf("%f ", mem[i]);
  }
  printf("\n");
}

// print_uint32 prints the 8 32-bit lanes.
void print_uint32(__m256i x)
{
  uint32_t mem[8];
  _mm256_storeu_si256((__m256i *)mem, x);
  for(int i = 0; i < 8; i++) {
    printf("0x%08x ", mem[i]);
  }
  printf("\n");
}

// encode_int encodes 4 (lat, lng) points.
void encode_int(double *points, uint64_t *output)
{
  // Quantize.
  __m256i q[2];
  for(int i = 0; i < 2; i++) {
    __m256i p = _mm256_loadu_pd(points + 4*i);
    p = _mm256_mul_pd(p, _mm256_set_pd(1/360.0, 1/180.0, 1/360.0, 1/180.0));
    p = _mm256_add_pd(p, _mm256_set1_pd(1.5));
    q[i] = _mm256_srli_epi64(_mm256_castpd_si256(p), 20);
  }

  // Spread
  
}

int main(int argc, char **argv)
{
  double point[8];
  uint64_t hash[4];

  assert(NUM_TEST_VECTORS == 4);
  for(int i = 0; i < 4; i++) {
    point[2*i] = test_vectors[i].lat;
    point[2*i+1] = test_vectors[i].lng;
  }

  encode_int(point, hash);

  for(int i = 0; i < 4; i++) {
    printf("hash[%d] = %016" PRIx64 "\texpect %016" PRIx64 "\n", i, hash[i], test_vectors[i].hash);
  }
}
