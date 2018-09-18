package tests

import (
	"github.com/Fengxq2014/coupon/common/cache"
	"testing"
	"time"
)

func set() {
	cache.GetCache().Set("test_key", "test_value", 5*time.Second)
}

func TestGet(t *testing.T) {
	set()
	v, err := cache.GetCache().Get("test_key")
	if err != nil {
		t.Error("Get value error ", err)
	}
	if "test_value" != v {
		t.Error("should get be test_value,but get ", v)
	}
}
