package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddBinary(t *testing.T) {
	result := addBinary("11", "1")
	require.Equal(t, result, "100")

	result = addBinary("1010", "1011")
	require.Equal(t, result, "10101")
}
