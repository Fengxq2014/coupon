package tests

import (
	"github.com/Fengxq2014/coupon/common/random"
	"testing"
)

func TestRandomStr(t *testing.T) {
	t.Log(random.RandStr(5))
}

func BenchmarkRandomStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.RandStr(5)
	}
}

func TestRandomNum(t *testing.T) {
	t.Log(random.RandRangeNum(1000, 9999))
}

func BenchmarkRandomNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.RandRangeNum(1000, 9999)
	}
}
