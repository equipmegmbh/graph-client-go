package graph

import (
	"encoding/json"
	"strconv"
	"strings"
)

// HexUid a struct which can represent an uid either as uint64 or as hex-string
type HexUid struct {
	Hex    string
	Uint64 uint64
	Valid  bool
}

func NewHexUid(hex string) *HexUid {
	result := &HexUid{}
	_ = result.SetByHex(hex)
	return result
}

func (h *HexUid) Clear() {
	h.Hex = ""
	h.Uint64 = 0
	h.Valid = false
}

func (h *HexUid) SetByHex(hex string) (err error) {
	// remove 0x suffix if found in the input string
	cleaned := strings.Replace(hex, "0x", "", -1)
	// base 16 for hexadecimal
	h.Uint64, err = strconv.ParseUint(cleaned, 16, 64)

	if err == nil {
		h.Hex = hex
		h.Valid = true
	} else {
		h.Hex = ""
		h.Valid = false
	}
	return
}

// MarshalJSON ...
func (h *HexUid) MarshalJSON() ([]byte, error) {
	if !h.Valid {
		return json.Marshal(nil)
	}
	buffer := make([]byte, 0)

	buffer = append(buffer, 0x22)
	buffer = append(buffer, []byte(h.Hex)...)
	buffer = append(buffer, 0x22)

	return buffer, nil
}

// UnmarshalJSON ...
func (h *HexUid) UnmarshalJSON(b []byte) error {
	h.Clear()

	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	var err error
	if b[0] == '"' && b[len(b)-1] == '"' {
		h.Hex = string(b[1 : len(b)-1])
		cleaned := strings.Replace(h.Hex, "0x", "", -1)
		h.Uint64, err = strconv.ParseUint(cleaned, 16, 64)
	} else {
		h.Uint64, err = strconv.ParseUint(string(b), 10, 64)
		if err == nil {
			h.Hex = "0x" + strconv.FormatUint(h.Uint64, 16)
		}
	}

	h.Valid = err == nil
	return nil
}
