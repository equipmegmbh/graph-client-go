package graph

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHexUid_SetByHex(t *testing.T) {
	testCases := []struct {
		name     string
		hex      string
		valid    bool
		expected string
		uint64   uint64
	}{
		{"ValidHex", "0x1a2b3c", true, "0x1a2b3c", 0x1a2b3c},
		{"InvalidHex", "invalid_hex", false, "", 0},
		{"EmptyHex", "", false, "", 0},
		{"HexWithoutPrefix", "1a2b3c", true, "1a2b3c", 0x1a2b3c},
		{"LargeHexValue", "0xFFFFFFFFFFFFFFFF", true, "0xFFFFFFFFFFFFFFFF", 0xFFFFFFFFFFFFFFFF},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uid := NewHexUid(tc.hex)
			require.Equal(t, tc.valid, uid.Valid)
			require.Equal(t, tc.expected, uid.Hex)
			require.Equal(t, tc.uint64, uid.Uint64)
		})
	}
}

func TestHexUid_Clear(t *testing.T) {
	testCases := []struct {
		name string
		hex  string
	}{
		{"ClearValidHex", "0x1a2b3c"},
		{"ClearEmptyHex", ""},
		{"ClearInvalidHex", "invalid_hex"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uid := NewHexUid(tc.hex)
			uid.Clear()
			require.False(t, uid.Valid)
			require.Equal(t, "", uid.Hex)
			require.Equal(t, uint64(0), uid.Uint64)
		})
	}
}

func TestHexUid_MarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		hex      string
		expected string
	}{
		{"MarshalValidHex", "0x1a2b3c", `"0x1a2b3c"`},
		{"MarshalEmptyHex", "invalid_hex", `null`},
		{"MarshalEmptyHex", "", `null`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			uid := NewHexUid(tc.hex)
			data, err := json.Marshal(uid)
			require.NoError(t, err)
			require.JSONEq(t, tc.expected, string(data))
		})
	}
}

func TestHexUid_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		json     string
		valid    bool
		expected string
		uint64   uint64
	}{
		{"UnmarshalValidHex", `"0x1a2b3c"`, true, "0x1a2b3c", 0x1a2b3c},
		{"UnmarshalValidHex", `"1a2b3c"`, true, "1a2b3c", 0x1a2b3c},
		{"UnmarshalInvalidHex", `"invalid_hex"`, false, "invalid_hex", 0},
		{"UnmarshalNumber", `1715004`, true, "0x1a2b3c", 1715004},
		{"UnmarshalNegativeNumber", `-1715004`, false, "", 0},
		{"UnmarshalNull", `null`, false, "", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var uid HexUid
			err := json.Unmarshal([]byte(tc.json), &uid)
			require.NoError(t, err)
			require.Equal(t, tc.valid, uid.Valid)
			require.Equal(t, tc.expected, uid.Hex)
			require.Equal(t, tc.uint64, uid.Uint64)
		})
	}
}
