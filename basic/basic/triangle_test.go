package basic

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{3, 4, 4},
		{3, 4, 3},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if res := CalTriangle(tt.a, tt.b); res != tt.c {
			t.Errorf("CalTriangle(%d, %d)=%d, but expect %d", tt.a, tt.b, res, tt.c)
		}
	}
}
