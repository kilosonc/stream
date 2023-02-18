package stream_test

import (
	"fmt"
	"testing"

	"github.com/kilosonc/stream"
	"github.com/stretchr/testify/assert"
)

func TestFill(t *testing.T) {
	count := 10
	intArr := make([]int, count)

	tmp := 0
	generator := func() int {
		t := tmp
		tmp++
		return t
	}

	stream.Fill(intArr, generator)
	assert.Equal(t, []int{0,1,2,3,4,5,6,7,8,9}, intArr)

	tmp = 0
	intArr = make([]int, count)
	stream.Fill(intArr[1: 5], generator)
	assert.Equal(t, []int{0,0,1,2,3,0,0,0,0,0}, intArr)

	type Cat struct {
		age int
	}

	tmp = 0
	catGenerator := func() *Cat {
		cat := Cat{
			age: tmp,
		}
		tmp++
		return &cat
	}
	catArr := make([]*Cat, 3)
	stream.Fill(catArr[0:2], catGenerator)
	assert.Equal(t, []*Cat{{age: 0}, {age: 1}, nil}, catArr)
}

func TestReduce(t *testing.T) {
	count := 3
	m := make(map[int]string, count)
	arrStr := make([]string, 0, count)
	
	for i := 0; i < count; i++ {
		arrStr = append(arrStr, fmt.Sprintf("reduce-%d", i))
	}
	
	accMap := func(m map[int]string,
		s string, index int) map[int]string {

		m[index] = s
		return m
	}
	m = stream.Reduce(arrStr,accMap, m)
	assert.Equal(t, map[int]string{0: "reduce-0", 1: "reduce-1", 2: "reduce-2"}, m)

	accInt := func(acc int, s int, index int) int {
		return acc + s
	}

	arrInt := make([]int, 0, 100)
	for i := 0;i < 100; i++ {
		arrInt = append(arrInt, i+1)
	}
	total := stream.Reduce(arrInt, accInt, 0)
	assert.Equal(t, 5050, total)
}