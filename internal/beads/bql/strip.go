package bql

import "strings"

// StripOrderBy removes the ORDER BY clause from a BQL query string.
// Since BQL grammar is strict (Filter → EXPAND → ORDER BY), ORDER BY
// is always the last clause, making it safe to truncate at that position.
//
// Returns the query without ORDER BY. If the query has no ORDER BY,
// returns it unchanged.
func StripOrderBy(query string) string {
	lexer := NewLexer(query)
	for {
		tok := lexer.NextToken()
		if tok.Type == TokenEOF {
			break
		}
		if tok.Type == TokenOrder {
			// Found ORDER keyword — strip everything from this position onward.
			// tok.Pos is 1-based (lexer increments pos before reading ch),
			// so subtract 1 to convert to 0-based string index.
			start := max(tok.Pos-1, 0)
			return strings.TrimRight(query[:start], " \t")
		}
	}
	return query
}

// HasOrderBy returns true if the query contains an ORDER BY clause.
func HasOrderBy(query string) bool {
	lexer := NewLexer(query)
	for {
		tok := lexer.NextToken()
		if tok.Type == TokenEOF {
			return false
		}
		if tok.Type == TokenOrder {
			return true
		}
	}
}
