package strings

import (
	"math"
)

// base represents the number whose power is computed for creating hash
// The greater the more chance for less collision.
//
// If s is the length of substring and b is the length of the source string
// Average time complexity of O(s + b).
// Worst case is O(sb).
const base = 128

type hashCache map[int]uint64

// Contains returns true if the substring exists in the source.
// It uses Rabin-Karp Substring algorithm to do so.
func Contains(source, substring string) bool {
	substrLength := len(substring)
	if substrLength > len(source) {
		return false
	}

	// cache is just to avoid repeated computations of base Powers
	var cache = make(hashCache)
	for i := 0; i < substrLength; i++ {
		cache[i] = uint64(math.Pow(base, float64(i)))
	}

	expectedHash := cache.compute(substring)
	computedHash := cache.compute(source[:substrLength])

	if computedHash == expectedHash {
		if substring == source[:substrLength] {
			return true
		}
	}

	highestPower := substrLength - 1
	for i := substrLength; i < len(source); i++ {
		computedHash = (computedHash-cache.code(source[i-substrLength], highestPower))*base + cache.code(source[i], 0)
		if computedHash == expectedHash {
			if substring == source[i-substrLength+1:i+1] {
				return true
			}
		}
	}

	return false
}

func (c hashCache) compute(str string) uint64 {
	var result uint64
	for i, j := len(str)-1, 0; i >= 0; i, j = i-1, j+1 {
		result += c.code(str[i], j)
	}

	return result
}

func (c hashCache) code(char byte, position int) uint64 {
	return c[position] * uint64(char)
}
