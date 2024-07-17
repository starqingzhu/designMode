/*
 * @Author: bsun
 * @Date: 2024-07-15 13:54:14
 * @Last Modified by: bsun
 * @Last Modified time: 2024-07-15 14:00:11
 */

package test

import (
	"designmode/behavioralDesign/interpreter"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInterpre(t *testing.T) {
	stats := map[string]float64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	tests := []struct {
		name  string
		stats map[string]float64
		rule  string
		want  bool
	}{
		{
			name:  "case1",
			stats: stats,
			rule:  "a > 1 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case2",
			stats: stats,
			rule:  "a < 2 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case3",
			stats: stats,
			rule:  "a < 5 && b > 1 && c < 10",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := interpreter.NewAlertRule(tt.rule)
			require.NoError(t, err)
			// assert.Equal(t, tt.want, r.Interpret(tt.stats))
		})
	}
}
