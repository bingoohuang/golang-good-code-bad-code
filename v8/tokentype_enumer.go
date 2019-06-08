package v8

import (
	"bytes"
	"fmt"
)

var _TokenTypeNameBytes = []byte("TITLEADEPALTNZADESARCIDARCTYPCEQPTMSGTXTCOMMENTEETFIRSPEEDESTDATAGEORTEPTS")

// {0, 5, 9, 14, 18, 23, 29, 34, 40, 47, 53, 58, 65, 68, 74}

// TITLE ADEP ALTNZ ADES ARCID ARCTYP CEQPT MSGTXT COMMENT EETFIR SPEED ESTDATA GEO RTEPTS
func TokenTypeBytes(s []byte) (TokenType, error) {
	if len(s) < 3 {
		return 0, fmt.Errorf("%s does not belong to TokenType values", s)
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

	return 0, fmt.Errorf("%s does not belong to TokenType values", s)
}
