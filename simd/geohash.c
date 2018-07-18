// SIMD implementation of integer geohash encoding.
//
// Based on work by Daniel Lemire and Geoff Langdale.
// https://lemire.me/blog/2018/01/09/how-fast-can-you-bit-interleave-32-bit-integers-simd-edition/

#include <inttypes.h>
#include <immintrin.h>

#include "geohash.h"

// spread the low 32 bits of each 64-bit lane into the even bit positions of
// the lane.
static inline __m256i spread(__m256i x)
{
  x = _mm256_and_si256(x, _mm256_set1_epi64x(0x00000000ffffffff));

  x = _mm256_or_si256(x, _mm256_slli_epi64(x, 16));
  x = _mm256_and_si256(x, _mm256_set1_epi64x(0x0000ffff0000ffff));

  x = _mm256_or_si256(x, _mm256_slli_epi64(x, 8));
  x = _mm256_and_si256(x, _mm256_set1_epi64x(0x00ff00ff00ff00ff));

  x = _mm256_or_si256(x, _mm256_slli_epi64(x, 4));
  x = _mm256_and_si256(x, _mm256_set1_epi64x(0x0f0f0f0f0f0f0f0f));

  x = _mm256_or_si256(x, _mm256_slli_epi64(x, 2));
  x = _mm256_and_si256(x, _mm256_set1_epi64x(0x3333333333333333));

  x = _mm256_or_si256(x, _mm256_slli_epi64(x, 1));
  x = _mm256_and_si256(x, _mm256_set1_epi64x(0x5555555555555555));

  return x;
}

void encode_int(double *lat, double *lng, uint64_t *output)
{
  // Quantize.
  __m256d latq = _mm256_loadu_pd(lat);
  latq = _mm256_mul_pd(latq, _mm256_set1_pd(1/180.0));
  latq = _mm256_add_pd(latq, _mm256_set1_pd(1.5));
  __m256i lati = _mm256_srli_epi64(_mm256_castpd_si256(latq), 20);

  __m256d lngq = _mm256_loadu_pd(lng);
  lngq = _mm256_mul_pd(lngq, _mm256_set1_pd(1/360.0));
  lngq = _mm256_add_pd(lngq, _mm256_set1_pd(1.5));
  __m256i lngi = _mm256_srli_epi64(_mm256_castpd_si256(lngq), 20);

  // Spread
  __m256i hash = _mm256_or_si256(spread(lati), _mm256_slli_epi64(spread(lngi), 1));
  _mm256_storeu_si256((__m256i *)output, hash);
}
