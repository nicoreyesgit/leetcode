package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidParentheses(t *testing.T) {
	testsCases := []struct {
		name   string
		input  string
		result bool
	}{
		{
			name:   "when all are closed should be ok",
			input:  "([])",
			result: true,
		},
		{
			name:   "when one is NOT closed should be NOK",
			input:  "(]",
			result: false,
		},
		{
			name:   "when closed are not ordered should return NOK",
			input:  "([)]",
			result: false,
		},
		{
			name:   "when closed are not ordered should return NOK",
			input:  ")(){}",
			result: false,
		},
		{
			name:   "when closed are not ordered should return NOK",
			input:  "[",
			result: false,
		},
	}

	for _, test := range testsCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.result, ValidParentheses(test.input))
		})
	}
}
