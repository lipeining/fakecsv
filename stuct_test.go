package fakecsv

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/stretchr/testify/assert"
)

func TestMakeStruct(t *testing.T) {
	gofakeit.Seed(0)
	s := MakeStruct(1, "xxxx", []int{}, []string{"a", "b"})
	gofakeit.Struct(s)
	assert.Equal(t, true, s != nil)
	fmt.Println(s)
	// 可以得到一个 fake 的结构体。说明只要符合 interface 的 struct 就可以使用 Struct 接口
	// &{-7986457512897628214 ibiye [] []}
}
