package good

import (
	"bytes"
	"fmt"
)

const _SubTokenTypeName = "PTIDETOFLGEOIDLATTDLONGTD"

var _SubTokenTypeNameBytes = []byte(_SubTokenTypeName)

var _SubTokenTypeIndex = [...]uint8{0, 4, 7, 9, 14, 19, 25}

func (i SubTokenType) String() string {
	if i < 0 || i >= SubTokenType(len(_SubTokenTypeIndex)-1) {
		return fmt.Sprintf("SubTokenType(%d)", i)
	}
	return _SubTokenTypeName[_SubTokenTypeIndex[i]:_SubTokenTypeIndex[i+1]]
}

var _SubTokenTypeValues = []SubTokenType{0, 1, 2, 3, 4, 5}

var _SubTokenTypeNameToValueMap = map[string]SubTokenType{
	_SubTokenTypeName[0:4]:   0,
	_SubTokenTypeName[4:7]:   1,
	_SubTokenTypeName[7:9]:   2,
	_SubTokenTypeName[9:14]:  3,
	_SubTokenTypeName[14:19]: 4,
	_SubTokenTypeName[19:25]: 5,
}

// PTID ETO FL GEOID LATTD LONGTD
// SubTokenTypeString retrieves an enum value from the enum constants string name.
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

// SubTokenTypeValues returns all values of the enum
func SubTokenTypeValues() []SubTokenType {
	return _SubTokenTypeValues
}

// IsASubTokenType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i SubTokenType) IsASubTokenType() bool {
	for _, v := range _SubTokenTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
