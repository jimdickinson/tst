package tst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPatternFullyConsumed(t *testing.T) {
	testCases := []struct {
		name          string
		pattern       string
		exactWildcard rune
		zeroOrMore    rune
		prefix        string
		expected      bool
	}{
		{
			name:          "Literal match",
			pattern:       "mango",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "One zeroOrMoreWildcardChar at the end",
			pattern:       "mango*",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "One zeroOrMoreWildcardChar at the beginning",
			pattern:       "*mango",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "Two zeroOrMoreWildcardChar characters, surrounding",
			pattern:       "*mango*",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "Two zeroOrMoreWildcardChar characters, surrounding a middle subset",
			pattern:       "*an*",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "Two zeroOrMoreWildcardChar characters, surrounding the last letter",
			pattern:       "*o*",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "Two zeroOrMoreWildcardChar characters, surrounding the first letter",
			pattern:       "*m*",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "Lots of zeroOrMoreWildcardChar characters",
			pattern:       "***ma*n*g*o*",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "ExactlyOneWildcardChar match",
			pattern:       "mang?",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "ExactlyOneWildcardChar in the right quantity",
			pattern:       "?????",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      true,
		},
		{
			name:          "Pattern with a repeating segment",
			pattern:       "ba*bad*",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "bababbababbad",
			expected:      true,
		},
		{
			name:          "Simplest literal does not match",
			pattern:       "x",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      false,
		},
		{
			name:          "Literal does not match",
			pattern:       "mango",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "supermango",
			expected:      false,
		},
		{
			name:          "Literal does not match the whole input",
			pattern:       "mango",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mangos",
			expected:      false,
		},
		{
			name:          "Too many exactly one wildchar chars",
			pattern:       "mang??",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      false,
		},
		{
			name:          "Exactly one wildcard chars with a different suffix",
			pattern:       "???ge",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      false,
		},
		{
			name:          "Way too many exactly one wildchar chars",
			pattern:       "??????",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mango",
			expected:      false,
		},
		{
			name:          "Mismatch with one zero or more character wildcard (1)",
			pattern:       "m*go",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "melon",
			expected:      false,
		},
		{
			name:          "Mismatch with one zero or more character wildcard (2)",
			pattern:       "*x",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "apple",
			expected:      false,
		},
		{
			name:          "Mismatch with two zero or more character wildcards (1)",
			pattern:       "m*g*oess",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "mangoes",
			expected:      false,
		},
		{
			name:          "Mismatch with two zero or more character wildcards (2)",
			pattern:       "*a*e",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "banana",
			expected:      false,
		},
		{
			name:          "Mismatch with three zero or more character wildcards",
			pattern:       "*m*go*z",
			exactWildcard: '?',
			zeroOrMore:    '*',
			prefix:        "supermangoes",
			expected:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pattern := NewWildcardPattern(tc.pattern, tc.exactWildcard, tc.zeroOrMore)
			result := isPatternFullyConsumed(pattern, tc.prefix)

			assert.Equal(t, tc.expected, result)
		})
	}
}
