package cosd

/*
#cgo LDFLAGS: -lm
#include <stdio.h>
#include <math.h>

double cosineSimilarity(float* A, float* B, int size) {
    float up = 0;
    float downl = 0;
    float downr = 0;
	int i;
    for (i = 0; i < size; i++) {
        float a = A[i];
        float b = B[i];
        up += a * b;
        downl += a * a;
        downr += b * b;
    }
    downl = sqrt(downl);
    downr = sqrt(downr);
    if (downl == 0 || downr == 0) {
        return 1;
    } else {
        return 1 - up / (downl * downr);
    }
}*/
import "C"
import (
	"unsafe"
)

func Cosd(a, b []float32) float32 {
	a2 := (*C.float)(unsafe.Pointer(&a[0]))
	b2 := (*C.float)(unsafe.Pointer(&b[0]))
	// Return (1-Similarity) as distance (lower distance means more similar)
	c := C.cosineSimilarity(a2, b2, C.int(len(a)))
	return float32(c)
}
