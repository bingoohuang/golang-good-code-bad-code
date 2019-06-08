package good

import (
	"bytes"
	"fmt"
)

const _TokenTypeName = "TITLEADEPALTNZADESARCIDARCTYPCEQPTMSGTXTCOMMENTEETFIRSPEEDESTDATAGEORTEPTS"
var _TokenTypeNameBytes = []byte(_TokenTypeName)

var _TokenTypeIndex = [...]uint8{0, 5, 9, 14, 18, 23, 29, 34, 40, 47, 53, 58, 65, 68, 74}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenTypeIndex)-1) {
		return fmt.Sprintf("TokenType(%d)", i)
	}
	return _TokenTypeName[_TokenTypeIndex[i]:_TokenTypeIndex[i+1]]
}

var _TokenTypeValues = []TokenType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

var _TokenTypeNameToValueMap = map[string]TokenType{
	_TokenTypeName[0:5]:   0,
	_TokenTypeName[5:9]:   1,
	_TokenTypeName[9:14]:  2,
	_TokenTypeName[14:18]: 3,
	_TokenTypeName[18:23]: 4,
	_TokenTypeName[23:29]: 5,
	_TokenTypeName[29:34]: 6,
	_TokenTypeName[34:40]: 7,
	_TokenTypeName[40:47]: 8,
	_TokenTypeName[47:53]: 9,
	_TokenTypeName[53:58]: 10,
	_TokenTypeName[58:65]: 11,
	_TokenTypeName[65:68]: 12,
	_TokenTypeName[68:74]: 13,
}

// TITLE ADEP ALTNZ ADES ARCID ARCTYP CEQPT MSGTXT COMMENT EETFIR SPEED ESTDATA GEO RTEPTS
func TokenTypeBytes(s []byte) (TokenType, error) {
	if len(s) < 3 {
		return  0, fmt.Errorf("%s does not belong to TokenType values", s)
	}

	switch s[0] {
	case 'T':
		if bytes.Equal(s, _TokenTypeNameBytes[0:5]) {
			return TokenTypeTITLE, nil
		}
	case 'A': // ADEP ALTNZ ADES ARCID ARCTYP
		if len(s) > 0 {
			switch s[3] {
			case 'P':
				if bytes.Equal(s, _TokenTypeNameBytes[5:9]) {
					return TokenTypeADEP, nil
				}
			case 'N':
				if bytes.Equal(s, _TokenTypeNameBytes[9:14]) {
					return TokenTypeALTNZ, nil
				}
			case 'S':
				if bytes.Equal(s, _TokenTypeNameBytes[14:18]) {
					return TokenTypeADES, nil
				}
			case 'I':
				if bytes.Equal(s, _TokenTypeNameBytes[18:23]) {
					return TokenTypeARCID, nil
				}
			case 'T':
				if bytes.Equal(s, _TokenTypeNameBytes[23:29]) {
					return TokenTypeARCTYP, nil
				}
			}
		}
	case 'C': // CEQPT MSGTXT COMMENT
		switch s[1] {
		case 'E':
			if bytes.Equal(s, _TokenTypeNameBytes[29:34]) {
				return TokenTypeCEQPT, nil
			}
		case 'O':
			if bytes.Equal(s, _TokenTypeNameBytes[40:47]) {
				return TokenTypeCOMMENT, nil
			}
		}
	case 'M': // MSGTXT COMMENT EETFIR SPEED ESTDATA GEO RTEPTS
		if bytes.Equal(s, _TokenTypeNameBytes[34:40]) {
			return TokenTypeMSGTXT, nil
		}
	case 'E': // EETFIR SPEED ESTDATA GEO RTEPTS
		switch s[1] {
		case 'E':
			if bytes.Equal(s, _TokenTypeNameBytes[47:53]) {
				return TokenTypeEETFIR, nil
			}
		case 'S':
			if bytes.Equal(s, _TokenTypeNameBytes[58:65]) {
				return TokenTypeESTDATA, nil
			}
		}
	case 'S': // SPEED
		if bytes.Equal(s, _TokenTypeNameBytes[53:58]) {
			return TokenTypeSPEED, nil
		}
	case 'G': // GEO
		if bytes.Equal(s, _TokenTypeNameBytes[65:68]) {
			return TokenTypeGEO, nil
		}
	case 'R': // RTEPTS
		if bytes.Equal(s, _TokenTypeNameBytes[68:74]) {
			return TokenTypeRTEPTS, nil
		}
	}

	return  0, fmt.Errorf("%s does not belong to TokenType values", s)
}

// TokenTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TokenTypeString(s string) (TokenType, error) {
	if val, ok := _TokenTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TokenType values", s)
}

// TokenTypeValues returns all values of the enum
func TokenTypeValues() []TokenType {
	return _TokenTypeValues
}

// IsATokenType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TokenType) IsATokenType() bool {
	for _, v := range _TokenTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
