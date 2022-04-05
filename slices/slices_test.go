package slices

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Student struct {
	Age int
}

func TestFilter(t *testing.T) {
	// 自动类型推导
	// Filter[int]([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(x, _ int) bool {
	//		return x%2 == 0
	//	})

	// 找出 2 的倍数
	filterIntSlice := Filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(x, _ int) bool {
		return x%2 == 0
	})
	assert.Equal(t, []int{2, 4, 6, 8}, filterIntSlice)

	// 找出包含 ja 的字符串
	filterStringSlice := Filter([]string{"jaronnie", "gocloudcoder", "nj-jay"}, func(s string, _ int) bool {
		return strings.Contains(s, "ja")
	})
	assert.Equal(t, []string{"jaronnie", "nj-jay"}, filterStringSlice)

	// 找出年龄大于等于 19 岁的学生
	filterStudentSlice := Filter([]Student{
		{18}, {19}, {20},
	}, func(student Student, _ int) bool {
		return student.Age >= 19
	})
	assert.Equal(t, 2, len(filterStudentSlice))
}
