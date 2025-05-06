package utils

// Ordered 代表所有可比大小排序的类型
type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}

func Min[T Ordered](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero // 如果没有输入值，返回类型的零值
	}

	min := values[0]
	for _, v := range values[1:] {
		if v < min { // 利用 comparable 约束进行直接比较
			min = v
		}
	}
	return min
}

func Max[T Ordered](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero // 如果没有输入值，返回类型的零值
	}

	min := values[0]
	for _, v := range values[1:] {
		if v > min { // 利用 comparable 约束进行直接比较
			min = v
		}
	}
	return min
}
