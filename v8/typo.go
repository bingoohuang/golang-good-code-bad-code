package v8

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
