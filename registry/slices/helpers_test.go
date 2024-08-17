package slices

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsComparable(t *testing.T) {
	r := NewRegistry()

	assert.True(t, r.isComparable(42))
	assert.True(t, r.isComparable(42.0))
	assert.True(t, r.isComparable("42"))
	assert.True(t, r.isComparable(true))
	assert.True(t, r.isComparable(`{"foo": "bar"}`))
	assert.True(t, r.isComparable("foo: bar"))
	assert.False(t, r.isComparable([]int{42}))
	assert.False(t, r.isComparable(map[string]int{"foo": 42}))
	assert.False(t, r.isComparable(struct{ Name string }{"example object"}))
	assert.False(t, r.isComparable(func() string { return "example function" }))
	assert.False(t, r.isComparable(fmt.Errorf("example error")))
	assert.False(t, r.isComparable(time.Now()))
	assert.False(t, r.isComparable(time.Second*5))
	assert.False(t, r.isComparable(make(chan any)))
	assert.False(t, r.isComparable(nil))

}

func TestSlicesRegistry_inList(t *testing.T) {
	r := NewRegistry()

	assert.True(t, r.inList([]any{1, 2, 3}, 2))
	assert.True(t, r.inList([]any{1, 2, 3}, 3))
	assert.False(t, r.inList([]any{1, 2, 3}, 4))
	assert.True(t, r.inList([]any{"a", "b", "c"}, "b"))
	assert.True(t, r.inList([]any{"a", "b", "c"}, "c"))
	assert.False(t, r.inList([]any{"a", "b", "c"}, "d"))
	assert.False(t, r.inList([]any{1, 2, 3}, 4.0))
	assert.True(t, r.inList([]any{"a", "b", "c"}, "b"))
	assert.True(t, r.inList([]any{"a", "b", "c"}, "c"))
	assert.False(t, r.inList([]any{"a", "b", "c"}, "d"))
	assert.True(t, r.inList([]any{"a", "b", "c"}, "b"))
	assert.True(t, r.inList([]any{"a", "b", "c"}, "c"))

	type testStruct struct{ Name string }
	test := []testStruct{{"a"}, {"b"}, {"c"}}
	var anyTest []any
	for _, item := range test {
		anyTest = append(anyTest, item)
	}
	assert.True(t, r.inList(anyTest, testStruct{"b"}))
	assert.True(t, r.inList(anyTest, testStruct{"c"}))
	assert.False(t, r.inList(anyTest, testStruct{"d"}))
}
