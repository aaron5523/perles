package bql

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStripOrderBy(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no ORDER BY",
			input: "type = bug",
			want:  "type = bug",
		},
		{
			name:  "filter with ORDER BY",
			input: "type = bug order by created desc",
			want:  "type = bug",
		},
		{
			name:  "ORDER BY only",
			input: "order by updated desc",
			want:  "",
		},
		{
			name:  "EXPAND with ORDER BY",
			input: "type = epic expand down depth 2 order by priority asc",
			want:  "type = epic expand down depth 2",
		},
		{
			name:  "filter EXPAND no ORDER BY",
			input: "type = bug expand down",
			want:  "type = bug expand down",
		},
		{
			name:  "case insensitive ORDER BY",
			input: "type = bug ORDER BY priority",
			want:  "type = bug",
		},
		{
			name:  "mixed case Order By",
			input: "type = bug Order By title asc",
			want:  "type = bug",
		},
		{
			name:  "multi-field ORDER BY",
			input: "status = open order by priority asc, created desc",
			want:  "status = open",
		},
		{
			name:  "empty query",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StripOrderBy(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHasOrderBy(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"no ORDER BY", "type = bug", false},
		{"with ORDER BY", "type = bug order by priority", true},
		{"ORDER BY only", "order by updated desc", true},
		{"case insensitive", "type = bug ORDER BY priority", true},
		{"empty", "", false},
		{"EXPAND only", "expand down", false},
		{"EXPAND + ORDER BY", "expand down order by priority", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasOrderBy(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}
