package gordle

import "strings"

// Hint describes the validity of a character in a word.
type Hint byte

const (
	unknownChar Hint = iota
	absentChar
	wrongPos
	correctPos
)

// String implementation the stringer interface.
func (h Hint) String() string {
	switch h {
	case absentChar:
		return "🔴"
	case wrongPos:
		return "🟡"
	case correctPos:
		return "💚"
	default:
		return "💔"
	}
}

// feedback is a list of hints, one per character of word.
type feedback []Hint

// String implementation of the stringer interface.
func (fb feedback) String() string {
	sb := strings.Builder{}
	for _, h := range fb {
		sb.WriteString(h.String())
	}
	return sb.String()
}

// Equal determines equality of two feedbacks.
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}

	for i, val := range fb {
		if val != other[i] {
			return false
		}
	}
	return true
}
