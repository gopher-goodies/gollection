package gollection_test

import (
	"testing"

	"github.com/azihsoyn/gollection"
	"github.com/stretchr/testify/assert"
)

func TestFlatMap(t *testing.T) {
	assert := assert.New(t)

	arr := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9, 10},
	}
	expect := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	res, err := gollection.New(arr).FlatMap(func(v int) int {
		return v * 2
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatMap_InterfaceSlice(t *testing.T) {
	assert := assert.New(t)
	arr := [][]interface{}{
		[]interface{}{1, 2, 3},
		[]interface{}{"a", "b"},
		nil,
	}
	expect := []interface{}{2, 4, 6, "a", "b"}

	res, err := gollection.New(arr).FlatMap(func(v interface{}) interface{} {
		if n, ok := v.(int); ok {
			return n * 2
		}
		return v
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatMap_EmptySlice(t *testing.T) {
	assert := assert.New(t)
	expect := []string{}
	res, err := gollection.New([][]int{}).FlatMap(func(v int) string {
		return ""
	}).Result()
	assert.NoError(err)
	assert.Equal(expect, res)
}

func TestFlatMap_NotSlice(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").FlatMap(func(v interface{}) interface{} {
		return ""
	}).Result()
	assert.Error(err)
}

func TestFlatMap_HavingError(t *testing.T) {
	assert := assert.New(t)
	_, err := gollection.New("not slice value").
		FlatMap(func(v interface{}) interface{} {
		return ""
	}).
		FlatMap(func(v interface{}) interface{} {
		return ""
	}).
		Result()
	assert.Error(err)
}
