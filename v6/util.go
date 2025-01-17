package v6

import (
	"bytes"
	"fmt"
	"strconv"
)

// Parse a line by returning the header (token name) and the value.
// Example: COMMENT TEST must returns COMMENT and TEST (in byte slices)
func parseLine(in []byte) ([]byte, []byte) {
	if len(in) == 0 {
		return nil, nil
	}

	i := bytes.IndexByte(in, ' ')

	if i == -1 {
		return in, nil
	}

	return in[:i], in[i+1:]
}

// As each flight level in represented as Fxxx (e.g. F350), this function simply parse the flight level to an int with an eventual error (if the int conversion fails for instance)
func extractFlightLevel(in string) (int, error) {
	fl, err := strconv.Atoi(in[1:])

	if err != nil {
		return -1, fmt.Errorf("flight level %v cannot be parsed", fl)
	}

	return fl, err
}
