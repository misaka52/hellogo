package main

import "testing"

func TestNoRepeating(t *testing.T) {
	tests := []struct {
		s   string
		len int
	}{
		{"abcba", 3},
		{"", 0},
		{"a", 1},
		{"aaa", 1},
		{"一二三二一", 3},
		//{"一二三二一", 1},
	}

	for _, tt := range tests {
		if res := MaxLenNoRepeatString(tt.s); res != tt.len {
			t.Errorf("MaxLenNoRepeatString(%s)=%d, but expect %d", tt.s, res, tt.len)
		}
	}
}

func BenchmarkMaxLenNoRepeatString(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	expect := 8
	for i := 0; i < 14; i++ {
		s = s + s
	}
	b.Logf("len(s)=%d", len(s))
	// 重置计时器
	b.ResetTimer()
	// b.N 表示系统决定运行次数
	for i := 0; i < b.N; i++ {
		if res := MaxLenNoRepeatString(s); res != expect {
			b.Errorf("MaxLenNoRepeatString(%s)=%d, but expect %d", s, res, expect)
		}
	}
}
