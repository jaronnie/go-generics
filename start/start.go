package start

type MyInt64 int64

type MyFloat64 float64

type Number1 interface {
	int64 | float64
}

type Number2 interface {
	~int64 | ~float64 // 表示包含自定义的别名
}

func AddNumber1[T Number1](ts ...T) T {
	var sum T
	for _, v := range ts {
		sum += v
	}
	return sum
}

func AddNumber2[T Number2](ts ...T) T {
	var sum T
	for _, v := range ts {
		sum += v
	}
	return sum
}
