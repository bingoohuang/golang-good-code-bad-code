package v6

import (
	"bytes"
	"fmt"
)

var _SubTokenTypeNameBytes = []byte("PTIDETOFLGEOIDLATTDLONGTD")

// PTID ETO FL GEOID LATTD LONGTD
// SubTokenTypeBytes retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func SubTokenTypeBytes(s []byte) (SubTokenType, error) {
	if len(s) < 2 {
		return 0, fmt.Errorf("%s does not belong to SubTokenType values", string(s))
	}

	switch s[0] {
	case 'P':
		if bytes.Equal(s, _SubTokenTypeNameBytes[0:4]) {
			return SubTokenTypePTID, nil
		}
	case 'E':
		if bytes.Equal(s, _SubTokenTypeNameBytes[4:7]) {
			return SubTokenTypeETO, nil
		}
	case 'F':
		if bytes.Equal(s, _SubTokenTypeNameBytes[7:9]) {
			return SubTokenTypeFL, nil
		}
	case 'G':
		if bytes.Equal(s, _SubTokenTypeNameBytes[9:14]) {
			return SubTokenTypeGEOID, nil
		}
	case 'L':
		switch s[1] {
		case 'A':
			if bytes.Equal(s, _SubTokenTypeNameBytes[14:19]) {
				return SubTokenTypeLATTD, nil
			}
		case 'O':
			if bytes.Equal(s, _SubTokenTypeNameBytes[19:25]) {
				return SubTokenTypeLONGTD, nil
			}
		}
	}

	return 0, fmt.Errorf("%s does not belong to SubTokenType values", string(s))
}
