package main

import (
	"fmt"
)

// computeLPS compute longest proper prefix
func computeLPS(s string) (lps []int) {
	n := len(s)
	if n == 0 {
		return []int{}
	}
	lps = make([]int, n)

	for i, k := 1, 0; i < n; i++ {
		for k > 0 && s[k] != s[i] {
			k = lps[k-1]
		}

		if s[k] == s[i] {
			k++
		}

		lps[i] = k
	}

	return
}

// kmpMatch do kmp match s against t
func kmpMatch(s string, t string) (matched []int) {
	if len(s) > len(t) {
		return []int{}
	}

	if len(s) == 0 { // -1 only means match everywhere
		return []int{-1}
	}

	// preallocate result
	matched = make([]int, 0, len(t)-len(s)+1)

	// compute LPS of S
	lps := computeLPS(s)

	n, m := len(t), len(s)
	for i, k := 0, 0; i < n; i++ {
		for k > 0 && s[k] != t[i] {
			k = lps[k-1]
		}

		if s[k] == t[i] {
			k++
		}

		if k == m {
			matched, k = append(matched, i-k+1), lps[k-1]
		}
	}

	return matched
}

func main() {
	fmt.Println(computeLPS("abbcabc"))
	fmt.Println(kmpMatch("aba", "aababbaba"))
}
