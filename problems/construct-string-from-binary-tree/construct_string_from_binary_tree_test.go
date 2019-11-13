package problem606

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected string
}

func TestTree2str(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 4},
			expected: "1(2(4))(3)",
		},
		{
			input:    []int{1, 2, 3, kit.NULL, 4},
			expected: "1(2()(4))(3)",
		},
	}
	for _, tc := range tests {
		output := tree2str(kit.SliceInt2TreeNode(tc.input))
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
