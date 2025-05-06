package utils

import (
	"reflect"
	"sort"
	"sync"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(0, 1, 1, 2)

	l := set.ToList()
	if len(l) != 3 {
		t.Errorf("set.ToList() len=%d, want 3", len(l))
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	if l[0] != 0 || l[1] != 1 || l[2] != 2 {
		t.Errorf("set.ToList()=%v, want [1,2,3]", l)
	}
}

func TestSet_Intersect(t *testing.T) {
	// 初始化 set1 并添加元素
	set1 := NewSet[int]()
	set1.Add(1, 2, 3)

	// 初始化 set2 并添加元素
	set2 := NewSet[int]()
	set2.Add(2, 3, 4)

	// 计算交集
	result := set1.Intersect(set2)

	// 验证结果是否正确
	expected := []int{2, 3}
	actual := result.ToList()

	if !reflect.DeepEqual(len(expected), len(actual)) {
		t.Errorf("Expected length %d, got %d", len(expected), len(actual))
	}

	for _, v := range expected {
		if !result.Exist(v) {
			t.Errorf("Expected value %v not found in result", v)
		}
	}
}

func TestSet_Union(t *testing.T) {
	// 初始化 set1 并添加元素
	set1 := NewSet[int]()
	set1.Add(1, 2, 3)

	// 初始化 set2 并添加元素
	set2 := NewSet[int]()
	set2.Add(3, 4, 5)

	// 计算并集
	result := set1.Union(set2)

	// 验证结果是否正确
	expected := []int{1, 2, 3, 4, 5}
	actual := result.ToList()

	if !reflect.DeepEqual(len(expected), len(actual)) {
		t.Errorf("Expected length %d, got %d", len(expected), len(actual))
	}

	for _, v := range expected {
		if !result.Exist(v) {
			t.Errorf("Expected value %v not found in result", v)
		}
	}
}

func TestSafeSet(t *testing.T) {
	set := NewSafeSet[int]()
	max := 100000
	wg := sync.WaitGroup{}
	wg.Add(3)
	fn := func() {
		for i := 0; i < max; i++ {
			set.Add(i)
		}
		wg.Done()
	}
	go fn()
	go fn()
	go fn()
	wg.Wait()

	l := set.ToList()
	if len(l) != max {
		t.Errorf("set.ToList() len=%d, want %d", len(l), max)
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
	if l[0] != 0 || l[1] != 1 || l[2] != 2 {
		t.Errorf("set.ToList(3)=%v, want [0,1,2]", l[:3])
	}
}

func TestSafeSet_Intersect(t *testing.T) {
	// 初始化 set1 并添加元素
	set1 := NewSafeSet[int]()
	set1.Add(1, 2, 3)

	// 初始化 set2 并添加元素
	set2 := NewSafeSet[int]()
	set2.Add(2, 3, 4)

	// 计算交集
	result := set1.Intersect(set2)

	// 验证结果是否正确
	expected := []int{2, 3}
	actual := result.ToList()

	if !reflect.DeepEqual(len(expected), len(actual)) {
		t.Errorf("Expected length %d, got %d", len(expected), len(actual))
	}

	for _, v := range expected {
		if !result.Exist(v) {
			t.Errorf("Expected value %v not found in result", v)
		}
	}
}

func TestSafeSet_Union(t *testing.T) {
	// 初始化 set1 并添加元素
	set1 := NewSafeSet[int]()
	set1.Add(1, 2, 3)

	// 初始化 set2 并添加元素
	set2 := NewSafeSet[int]()
	set2.Add(3, 4, 5)

	// 计算并集
	result := set1.Union(set2)

	// 验证结果是否正确
	expected := []int{1, 2, 3, 4, 5}
	actual := result.ToList()

	if !reflect.DeepEqual(len(expected), len(actual)) {
		t.Errorf("Expected length %d, got %d", len(expected), len(actual))
	}

	for _, v := range expected {
		if !result.Exist(v) {
			t.Errorf("Expected value %v not found in result", v)
		}
	}
}
