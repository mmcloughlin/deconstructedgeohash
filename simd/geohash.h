#ifndef GEOHASH_H_
#define GEOHASH_H_

#include <inttypes.h>

#define BATCH_SIZE (4)

// encode_int encodes 4 (lat, lng) points.
void encode_int(double *lat, double *lng, uint64_t *output);

#endif // GEOHASH_H_
