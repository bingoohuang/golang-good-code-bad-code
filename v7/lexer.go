package v7

type lexeme []byte

func (x lexeme) isNewline() bool {
	return len(x) == 1 && x[0] == '\n'
}

func (x lexeme) isCommand() bool {
	return len(x) >= 2 && x[0] == '-'
}

const (
	start = iota
	word
)

var newLine = []byte("\n")

func lex(src []byte) (words []lexeme) {
	l := len(src)
	words = make([]lexeme, 0, l/4)

	state := start
	// commandType := noCommand
	i := 0

	for j := 0; ; j++ {
		if j == l {
			if i != j {
				words = append(words, src[i:j])
			}
			break
		}
		switch state {

		case start:
			switch src[j] {
			case ' ', '\t', '\n':
				// Ignore
			default:
				state = word
			}
			i = j

		case word:
			switch src[j] {
			case ' ', '\t':
				words = append(words, src[i:j])
				state = start
			case '\n':
				words = append(words, src[i:j])
				words = append(words, newLine)
				state = start
			default:
				// Keeping reading the word
			}
		}
	}
	return words
}
