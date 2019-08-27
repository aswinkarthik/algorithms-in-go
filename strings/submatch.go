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

// Submatch returns true if the substring exists in the source.
// It uses Rabin-Karp Substring algorithm to do so.
func Submatch(source, substring string) bool {
	expectedHash := hash(substring)
	substrLength := len(substring)
	if substrLength > len(source) {
		return false
	}

	computedHash := hash(source[:substrLength])

	if computedHash == expectedHash {
		return true
	}

	for i := substrLength; i < len(source); i++ {
		computedHash = (computedHash-code(source[i-substrLength], substrLength-1))*base + code(source[i], 0)
		if computedHash == expectedHash {
			return true
		}
	}

	return false
}

// powerHashCache is just to avoid repeated computations of base Powers
var powerHashCache = make(map[int]uint64)

func hash(str string) uint64 {
	var result uint64
	for i, j := len(str)-1, 0; i >= 0; i, j = i-1, j+1 {
		result += code(str[i], j)
	}

	return result
}

func code(char byte, position int) uint64 {
	computed, present := powerHashCache[position]
	if !present {
		computed = uint64(math.Pow(base, float64(position)))
		powerHashCache[position] = computed
	}
	return computed * uint64(char)
}
