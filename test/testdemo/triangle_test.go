package testdemo

import (
	"golangdemo/test/basic"
	"testing"
)

//func TestTriangle(t *testing.T) {
//
//	tests := []struct{ a, b, c int }{
//		{3, 4, 5},
//		{5, 12, 13},
//		{8, 15, 17},
//		{111, 222, 333},
//		{33333, 44444, 55555},
//		//{99999, 12121212, 15151515},
//	}
//
//	for a, tt := range tests {
//		fmt.Println(a, tt)
//		if triangle := basic.Triangle(tt.b, tt.a); triangle != tt.c {
//			t.Errorf("calcTriangle(%d, %d); got %d, excepted %d .\n", tt.a, tt.b, triangle, tt.c)
//		}
//	}
//}
//
//func BenchmarkTriangle(bb *testing.B) {
//	var a, b, c = 3333, 4444, 5555
//
//	bb.ResetTimer()
//	for i := 0; i < bb.N; i++ {
//		if triangle := basic.Triangle(b, a); triangle != c {
//			bb.Errorf("calcTriangle(%d, %d); got %d, excepted %d .\n", a, b, triangle, c)
//		}
//	}
//
//}

func BenchmarkSubLen(bb *testing.B) {

	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑回飞花"
	for i := 0; i < 10; i++ {
		s = s + s
	}
	c := 8
	bb.Logf("len(s) = %d \n", len(s))

	bb.ResetTimer()
	for i := 0; i < bb.N; i++ {
		if triangle := basic.GetNonRepeatLen(s); triangle != c {
			bb.Errorf("GetNonRepeatLen(%s); got %d, excepted %d .\n", s, triangle, c)
		}
	}

}
