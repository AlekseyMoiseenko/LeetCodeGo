package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  [][]string
	}{
		{
			name:  "классический пример",
			input: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name:  "пустые строки",
			input: []string{"", ""},
			want:  [][]string{{"", ""}},
		},
		{
			name:  "одна строка",
			input: []string{"a"},
			want:  [][]string{{"a"}},
		},
		{
			name:  "несколько групп разной длины",
			input: []string{"ab", "ba", "abc", "cba", "bac"},
			want: [][]string{
				{"ab", "ba"},
				{"abc", "cba", "bac"},
			},
		},
		{
			name:  "одинаковые строки",
			input: []string{"aa", "aa", "aa"},
			want:  [][]string{{"aa", "aa", "aa"}},
		},
		{
			name:  "строки максимальной длины",
			input: []string{strings.Repeat("a", 100), strings.Repeat("a", 100)},
			want: [][]string{
				{strings.Repeat("a", 100), strings.Repeat("a", 100)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.input)
			require.Equal(t, normalize(tt.want), normalize(got))
		})
	}
}

func normalize(groups [][]string) [][]string {
	res := make([][]string, len(groups))
	for i, g := range groups {
		res[i] = slices.Clone(g)
		slices.Sort(res[i])
	}
	slices.SortFunc(res, slices.Compare)
	return res
}
