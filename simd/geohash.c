// SIMD implementation of integer geohash encoding.
//
// Based on work by Daniel Lemire and Geoff Langdale.
// https://lemire.me/blog/2018/01/09/how-fast-can-you-bit-interleave-32-bit-integers-simd-edition/

#include <inttypes.h>
#include <immintrin.h>

#include "geohash.h"

#ifdef IACA
#include "iacaMarks.h"
#define KERNEL_START IACA_START
#define KERNEL_END IACA_END
#else
#define KERNEL_START
#define KERNEL_END
#endif

// spread the low 32 bits of each 64-bit lane into the even bit positions of
// the lane.
static inline __m256i spread(__m256i x)
{
  x  = _mm256_shuffle_epi8(x, _mm256_set_epi8(
            -1, 11, -1, 10, -1, 9, -1, 8,
            -1,  3, -1,  2, -1, 1, -1, 0,
            -1, 11, -1, 10, -1, 9, -1, 8,
            -1,  3, -1,  2, -1, 1, -1, 0));

  const __m256i lut = _mm256_set_epi8(
            85, 84, 81, 80, 69, 68, 65, 64,
            21, 20, 17, 16,  5,  4,  1,  0,
            85, 84, 81, 80, 69, 68, 65, 64,
            21, 20, 17, 16,  5,  4,  1,  0);

  __m256i lo = _mm256_shuffle_epi8(lut, _mm256_and_si256(x, _mm256_set1_epi8(0xf)));

  __m256i hi = _mm256_and_si256(x, _mm256_set1_epi8(0xf0));
  hi = _mm256_shuffle_epi8(lut, _mm256_srli_epi64(hi, 4));

  x = _mm256_or_si256(lo, _mm256_slli_epi64(hi, 8));

  return x;
}

void encode_int(double *lat, double *lng, uint64_t *output)
{
  KERNEL_START

  // Quantize.
  __m256d latq = _mm256_loadu_pd(lat);
  latq = _mm256_mul_pd(latq, _mm256_set1_pd(1/180.0));
  latq = _mm256_add_pd(latq, _mm256_set1_pd(1.5));
  __m256i lati = _mm256_srli_epi64(_mm256_castpd_si256(latq), 20);

  __m256d lngq = _mm256_loadu_pd(lng);
  lngq = _mm256_mul_pd(lngq, _mm256_set1_pd(1/360.0));
  lngq = _mm256_add_pd(lngq, _mm256_set1_pd(1.5));
  __m256i lngi = _mm256_srli_epi64(_mm256_castpd_si256(lngq), 20);

  // Spread.
  __m256i hash = _mm256_or_si256(spread(lati), _mm256_slli_epi64(spread(lngi), 1));
  _mm256_storeu_si256((__m256i *)output, hash);

  KERNEL_END
}
