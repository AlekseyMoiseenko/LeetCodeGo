package hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func board(rows ...string) [][]byte {
	b := make([][]byte, len(rows))
	for i, row := range rows {
		b[i] = []byte(row)
	}
	return b
}

// Constraints:
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.
func TestIsValidSudoku(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		board    [][]byte
		expected bool
	}{
		{
			name: "leetcode example 1 - valid",
			board: board(
				"53..7....",
				"6..195...",
				".98....6.",
				"8...6...3",
				"4..8.3..1",
				"7...2...6",
				".6....28.",
				"...419..5",
				"....8..79",
			),
			expected: true,
		},
		{
			name: "leetcode example 2 - duplicate 8 in top-left box",
			board: board(
				"83..7....",
				"6..195...",
				".98....6.",
				"8...6...3",
				"4..8.3..1",
				"7...2...6",
				".6....28.",
				"...419..5",
				"....8..79",
			),
			expected: false,
		},
		{
			name: "empty board is valid",
			board: board(
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			),
			expected: true,
		},
		{
			name: "duplicate in a row",
			board: board(
				"55.......",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			),
			expected: false,
		},
		{
			name: "duplicate in a column, rows not adjacent",
			board: board(
				"5........",
				".........",
				".........",
				".........",
				"5........",
				".........",
				".........",
				".........",
				".........",
			),
			expected: false,
		},
		{
			name: "duplicate in a 3x3 box but different row and column",
			board: board(
				"5........",
				".........",
				"..5......",
				".........",
				".........",
				".........",
				".........",
				".........",
				".........",
			),
			expected: false,
		},
		{
			name: "single fully valid completed board",
			board: board(
				"534678912",
				"672195348",
				"198342567",
				"859761423",
				"426853791",
				"713924856",
				"961537284",
				"287419635",
				"345286179",
			),
			expected: true,
		},
		{
			name: "completed board with one row broken",
			board: board(
				"534678912",
				"672195348",
				"198342567",
				"859761423",
				"426853791",
				"713924856",
				"961537284",
				"287419635",
				"345286171", // last cell duplicates the '1' already in this row
			),
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := isValidSudoku(tc.board)
			require.Equal(t, tc.expected, got)
		})
	}
}
