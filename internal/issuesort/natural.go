// Package issuesort provides natural sorting for beads issue IDs.
package issuesort

import (
	"sort"
	"strconv"
	"strings"
)

// NaturalIDs sorts issue IDs so that numeric segments are compared
// numerically rather than lexicographically.
// e.g., "x.1.2" < "x.1.10" instead of "x.1.10" < "x.1.2".
func NaturalIDs(ids []string) {
	if len(ids) <= 1 {
		return
	}
	sort.Slice(ids, func(i, j int) bool {
		return naturalLess(ids[i], ids[j])
	})
}

// naturalLess compares two strings with natural ordering by splitting on "."
// and comparing numeric segments as integers.
func naturalLess(a, b string) bool {
	aParts := strings.Split(a, ".")
	bParts := strings.Split(b, ".")

	for k := 0; k < len(aParts) && k < len(bParts); k++ {
		if aParts[k] == bParts[k] {
			continue
		}
		aNum, aErr := strconv.Atoi(aParts[k])
		bNum, bErr := strconv.Atoi(bParts[k])
		if aErr == nil && bErr == nil {
			return aNum < bNum
		}
		return aParts[k] < bParts[k]
	}
	return len(aParts) < len(bParts)
}
