package filter

import (
	"crypto/sha256"
)

func expensiveScore(input []byte) int {
	sum := input

	// symulacja ciężkiego targetingu / ML feature hashing
	for i := 0; i < 200; i++ {
		h := sha256.Sum256(sum)
		sum = h[:]
	}

	return int(sum[0])
}
