package good

import (
	"bytes"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type TokenType int

const (
	// List of ADEXP tokens
	TokenTypeTITLE TokenType = iota
	TokenTypeADEP
	TokenTypeALTNZ
	TokenTypeADES
	TokenTypeARCID
	TokenTypeARCTYP
	TokenTypeCEQPT
	TokenTypeMSGTXT
	TokenTypeCOMMENT
	TokenTypeEETFIR
	TokenTypeSPEED
	TokenTypeESTDATA
	TokenTypeGEO
	TokenTypeRTEPTS
)

type SubTokenType int

const (
	// List of ADEXP subtokens
	SubTokenTypePTID SubTokenType = iota
	SubTokenTypeETO
	SubTokenTypeFL
	SubTokenTypeGEOID
	SubTokenTypeLATTD
	SubTokenTypeLONGTD
)

var (
	bytesNewline     = []byte("\n")
	bytesNewlineDash = []byte("\n-")
	bytesBegin       = []byte("-BEGIN")
	bytesEnd         = []byte("-END")
	bytesEmpty       = []byte(" ")
	bytesComment     = []byte("//")
	bytesDash        = []byte("-")

	// Map containing the mapping function given a specific token name
	factory = map[TokenType]func([]byte) MyToken{
		TokenTypeTITLE:   parseSimpleToken(TokenTypeTITLE),
		TokenTypeADEP:    parseSimpleToken(TokenTypeADEP),
		TokenTypeALTNZ:   parseSimpleToken(TokenTypeALTNZ),
		TokenTypeADES:    parseSimpleToken(TokenTypeADES),
		TokenTypeARCID:   parseSimpleToken(TokenTypeARCID),
		TokenTypeARCTYP:  parseSimpleToken(TokenTypeARCTYP),
		TokenTypeCEQPT:   parseSimpleToken(TokenTypeCEQPT),
		TokenTypeMSGTXT:  parseSimpleToken(TokenTypeMSGTXT),
		TokenTypeCOMMENT: parseSimpleToken(TokenTypeCOMMENT),
		TokenTypeEETFIR:  parseSimpleToken(TokenTypeEETFIR),
		TokenTypeSPEED:   parseSimpleToken(TokenTypeSPEED),
		TokenTypeESTDATA: parseComplexToken(TokenTypeESTDATA),
		TokenTypeGEO:     parseComplexToken(TokenTypeGEO),
		TokenTypeRTEPTS:  parseComplexToken(TokenTypeRTEPTS),
	}
)

// Parse an ADEXP message using a byte list as an input. This function returns a Message and an eventual error in case of a parsing error.
func ParseAdexpMessage(bytes []byte) (Message, error) {
	log.Debugf("parsing: %v", string(bytes))

	// Preprocessing
	preprocessed, err := preprocess(bytes)
	if err != nil {
		return Message{}, err
	}

	// Actual processing
	message, err := process(preprocessed)
	if err != nil {
		return Message{}, err
	}

	log.Debugf("returning message: %v", message)

	return message, nil
}

// Preprocessing of an ADEXP message (cleaning white spaces, rearranging multi-lined tokens etc.). Returns a byte list cleansed and an eventual error if the input is invalid.
func preprocess(in []byte) ([][]byte, error) {
	if len(in) == 0 {
		log.Errorf("input is empty")
		return nil, errors.New("input is empty")
	}

	lines := bytes.Split(in, bytesNewline)
	var result [][]byte
	var currentLine []byte

	for _, line := range lines {
		if bytes.HasPrefix(line, bytesEnd) {
			// Nothing
		} else if bytes.HasPrefix(line, bytesBegin) {
			result = append(result, currentLine)

			trimed := trim(line)
			currentLine = append(bytesDash, trimed[len(bytesBegin)+1:]...)
		} else if bytes.HasPrefix(line, bytesDash) {
			result = append(result, currentLine)

			currentLine = trim(line)
		} else if bytes.HasPrefix(line, bytesEmpty) {
			currentLine = append(append(currentLine, bytesEmpty...), trim(line)...)
		} else {
			currentLine = append(append(currentLine, bytesEmpty...), trim(line)...)
		}
	}

	if len(currentLine) > 0 {
		result = append(result, currentLine)
	}

	return result, nil
}

var myTokenMsgFill = map[TokenType]func(st MyToken, msg *Message) error{
	TokenTypeTITLE:   func(st MyToken, msg *Message) error { msg.Title = st.simple; return nil },
	TokenTypeADEP:    func(st MyToken, msg *Message) error { msg.Adep = st.simple; return nil },
	TokenTypeALTNZ:   func(st MyToken, msg *Message) error { msg.Alternate = st.simple; return nil },
	TokenTypeADES:    func(st MyToken, msg *Message) error { msg.Ades = st.simple; return nil },
	TokenTypeARCID:   func(st MyToken, msg *Message) error { msg.Arcid = st.simple; return nil },
	TokenTypeARCTYP:  func(st MyToken, msg *Message) error { msg.ArcType = st.simple; return nil },
	TokenTypeCEQPT:   func(st MyToken, msg *Message) error { msg.Ceqpt = st.simple; return nil },
	TokenTypeMSGTXT:  func(st MyToken, msg *Message) error { msg.MessageText = st.simple; return nil },
	TokenTypeCOMMENT: func(st MyToken, msg *Message) error { msg.Comment = st.simple; return nil },
	TokenTypeEETFIR:  func(st MyToken, msg *Message) error { msg.Eetfir = append(msg.Eetfir, st.simple); return nil },
	TokenTypeSPEED:   func(st MyToken, msg *Message) error { msg.Speed = append(msg.Speed, st.simple); return nil },

	TokenTypeESTDATA: func(st MyToken, msg *Message) error {
		for _, v := range st.complex {
			fl, err := extractFlightLevel(v[SubTokenTypeFL])
			if err != nil {
				return fmt.Errorf("flight level %v cannot be parsed", fl)
			}
			msg.Estdata = append(msg.Estdata, estdata{v[SubTokenTypePTID], v[SubTokenTypeETO], fl})
		}
		return nil
	},
	TokenTypeGEO: func(st MyToken, msg *Message) error {
		for _, v := range st.complex {
			msg.Geo = append(msg.Geo, geo{v[SubTokenTypeGEOID], v[SubTokenTypeLATTD], v[SubTokenTypeLONGTD]})
		}
		return nil
	},
	TokenTypeRTEPTS: func(st MyToken, msg *Message) error {
		for _, v := range st.complex {
			fl, err := extractFlightLevel(v[SubTokenTypeFL])
			if err != nil {
				return fmt.Errorf("flight level %v cannot be parsed", fl)
			}
			msg.RoutePoints = append(msg.RoutePoints, rtepts{v[SubTokenTypePTID], fl, v[SubTokenTypeETO]})
		}
		return nil
	},
}

// Processing of an ADEXP message. Returns a Message structure and an eventual error in case of a processing error.
func process(in [][]byte) (Message, error) {
	msg := Message{Type: AdexpType}

	// Gather the goroutine results
	for _, line := range in {
		data, ok := mapLine(line)
		// A mapper function can return a nil value (a line is potentially invalid, a comment etc.). In that case we simply discard the line.
		if !ok {
			continue
		}

		// Enrich the message depending on the data type sent by the goroutines
		if fill, ok := myTokenMsgFill[data.token]; ok {
			if err := fill(data, &msg); err != nil {
				return Message{}, err
			}
		} else {
			log.Errorf("unexpected token type %v", data.token)
			return Message{}, fmt.Errorf("unexpected token type %v", data.token)
		}
	}

	return msg, nil
}

// Process a line and returns a token
func mapLine(in []byte) (MyToken, bool) {
	// Filter empty lines and comment lines
	if len(in) == 0 || bytes.HasPrefix(in, bytesComment) {
		return MyToken{}, false
	}
	if in[0] != '-' {
		log.Warnf("Line doesn't start with a proper token: %q", string(in))
		return MyToken{}, false
	}

	token, value := parseLine(in[1:])
	if token == nil {
		log.Warnf("Token name is empty on line %v", string(in))
		return MyToken{}, false
	}

	sToken, err := TokenTypeBytes(token)
	if err != nil {
		log.Warnf("unknown TokenType %v", string(token))
		return MyToken{}, false
	}

	// Checks in the factory map if the token has been configured
	if f, contains := factory[sToken]; contains {
		return f(value), true
	}

	log.Warnf("Token %v is not managed by the parser", string(in))
	return MyToken{}, false
}

// Parse a simple token and returns a simpleToken structure
func parseSimpleToken(token TokenType) func(value []byte) MyToken {
	return func(value []byte) MyToken {
		return MyToken{token: token, simple: string(value)}
	}
}

// Parse a complex token and returns a commplexToken structure
func parseComplexToken(token TokenType) func(value []byte) MyToken {
	return func(value []byte) MyToken {
		if len(value) == 0 {
			log.Warnf("Empty value")
			return MyToken{token: token, complex: nil}
		}

		var v []map[SubTokenType]string
		currentMap := make(map[SubTokenType]string)

		// Find all subfields
		matches := findSubfields(value)

		// Iterate over each subfields to enrich the returned data
		for _, sub := range matches {
			h, l := parseLine(sub)

			subTokenType, err := SubTokenTypeBytes(h)
			if err != nil {
				continue
			}
			if _, contains := currentMap[subTokenType]; contains {
				v = append(v, currentMap)
				currentMap = make(map[SubTokenType]string)
			}

			currentMap[subTokenType] = string(trim(l))
		}

		// Append the latest map
		v = append(v, currentMap)

		return MyToken{token: token, complex: v}
	}
}

// Extract subfields from a line (with dash removed).
// E.g.
//       "-ESTDATA -PTID XETBO -ETO 170302032300 -FL F390"
//    -> ["ESTDATA ", "PTID XETBO ", "ETO 170302032300 ", "FL F390"]
// This is efficient because each element is a subslice of the original line.
func findSubfields(value []byte) (subfields [][]byte) {
	subfields = bytes.Split(value, bytesDash)
	if len(subfields) > 0 && len(trim(subfields[0])) == 0 {
		subfields = subfields[1:]
	}
	return subfields
}

// This custom loop is faster than the generic-purpose bytes.Trim .
// It expects 1 char == 1 byte (no multi-byte UTF-8 runes)
func trim(s []byte) []byte {
	const space = ' '
	n := len(s)
	low, high := 0, n
	for low < n && s[low] == space {
		low++
	}
	for high > low && s[high-1] == space {
		high--
	}
	return s[low:high]
}
