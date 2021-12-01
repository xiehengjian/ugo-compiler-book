// Code generated by "stringer -type=TokenType"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EOF-0]
	_ = x[ERROR-1]
	_ = x[IDENT-2]
	_ = x[NUMBER-3]
	_ = x[PACKAGE-4]
	_ = x[FUNC-5]
	_ = x[ADD-6]
	_ = x[SUB-7]
	_ = x[MUL-8]
	_ = x[DIV-9]
	_ = x[LPAREN-10]
	_ = x[RPAREN-11]
	_ = x[LBRACE-12]
	_ = x[RBRACE-13]
	_ = x[SEMICOLON-14]
}

const _TokenType_name = "EOFERRORIDENTNUMBERPACKAGEFUNCADDSUBMULDIVLPARENRPARENLBRACERBRACESEMICOLON"

var _TokenType_index = [...]uint8{0, 3, 8, 13, 19, 26, 30, 33, 36, 39, 42, 48, 54, 60, 66, 75}

func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}