package tst

import "strings"

type WildcardPattern struct {
	pattern                   []rune
	exactlyOneWildcardChar    rune
	zeroOrMoreWildcardChar    rune
	hasZeroOrMoreWildcardChar bool
}

func NewWildcardPattern(pattern string, exactlyOneWildcardChar rune, zeroOrMoreWildcardChar rune) *WildcardPattern {
	// Check if the pattern contains zeroOrMoreWildcardChar and attempt an early return
	patternRunes := []rune(pattern)

	hasZeroOrMoreWildcardChar := false
	for _, ch := range patternRunes {
		if ch == zeroOrMoreWildcardChar {
			hasZeroOrMoreWildcardChar = true
			break
		}
	}

	return &WildcardPattern{
		pattern:                   patternRunes,
		exactlyOneWildcardChar:    exactlyOneWildcardChar,
		zeroOrMoreWildcardChar:    zeroOrMoreWildcardChar,
		hasZeroOrMoreWildcardChar: hasZeroOrMoreWildcardChar,
	}
}

func (tst *TernarySearchTrie) SearchWildcard(wp *WildcardPattern) []string {
	results := make([]string, 0)
	tst.searchWildcard(tst.root, wp, "", 0, &results)
	return results
}

func (tst *TernarySearchTrie) searchWildcard(node *Node, wp *WildcardPattern, currentPrefix string, index int, results *[]string) {
	if node == nil {
		return
	}

	tst.searchWildcard(node.left, wp, currentPrefix, index, results)

	newPrefix := currentPrefix + string(node.char)
	if canMatch(wp, newPrefix) {
		if node.value != 0 && isPatternFullyConsumed(wp, newPrefix) {
			*results = append(*results, newPrefix)
		}
		tst.searchWildcard(node.middle, wp, newPrefix, index+1, results)
	}

	tst.searchWildcard(node.right, wp, currentPrefix, index, results)
}

func canMatch(wp *WildcardPattern, prefix string) bool {
	prefixRunes := []rune(prefix)

	for i := 0; i < len(prefixRunes) && i < len(wp.pattern); i++ {
		if wp.pattern[i] == wp.zeroOrMoreWildcardChar {
			return true
		}
		if wp.pattern[i] != wp.exactlyOneWildcardChar &&
			wp.pattern[i] != prefixRunes[i] {
			return false
		}
	}
	return true
}

func isPatternFullyConsumed(wp *WildcardPattern, prefix string) bool {
	prefixRunes := []rune(prefix)

	if !wp.hasZeroOrMoreWildcardChar {
		if len(wp.pattern) != len(prefixRunes) {
			return false
		}
		for i, ch := range wp.pattern {
			if ch != wp.exactlyOneWildcardChar && ch != prefixRunes[i] {
				return false
			}
		}
		return true
	}

	segments := strings.Split(string(wp.pattern), string(wp.zeroOrMoreWildcardChar))
	filteredSegments := make([]string, 0, len(segments))
	for _, segment := range segments {
		if len(segment) > 0 {
			filteredSegments = append(filteredSegments, segment)
		}
	}
	segments = filteredSegments

	currentPrefixIndex := 0
	for _, segment := range segments {
		segmentRunes := []rune(segment)

		segmentIndex := 0
		for currentPrefixIndex < len(prefixRunes) {
			if segmentRunes[segmentIndex] == wp.exactlyOneWildcardChar || segmentRunes[segmentIndex] == prefixRunes[currentPrefixIndex] {
				segmentIndex++
				if segmentIndex == len(segmentRunes) {
					break
				}
			} else if segmentIndex > 0 {
				segmentIndex = 0
			}
			currentPrefixIndex++
		}

		if segmentIndex != len(segmentRunes) {
			return false
		}
	}

	//return currentPrefixIndex == len(prefixRunes)
	return true
}
