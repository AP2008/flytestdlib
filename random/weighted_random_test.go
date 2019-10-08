package random

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	key string
	val int
}

func (t testData) Compare(to Comparable) bool {
	if strings.Contains(t.key, "sort") {
		return t.key < to.(testData).key
	}
	return t.val < to.(testData).val
}

func TestDeterministicWeightedRandomStr(t *testing.T) {
	item1 := testData{
		key: "sort_key1",
		val: 1,
	}
	item2 := testData{
		key: "sort_key2",
		val: 2,
	}
	entries := []Entry{
		{
			Item:   item1,
			Weight: 0.4,
		},
		{
			Item:   item2,
			Weight: 0.6,
		},
	}
	rand, err := NewWeightedRandom(entries)
	assert.Nil(t, err)
	retItem, err := rand.GetWithSeed("ab")
	assert.Nil(t, err)
	assert.Equal(t, item1, retItem)

	assert.Nil(t, err)
	for i := 1; i <= 10; i++ {
		retItem, err := rand.GetWithSeed("hi")
		assert.Nil(t, err)
		assert.Equal(t, item2, retItem)
	}
}

func TestDeterministicWeightedRandomInt(t *testing.T) {
	item1 := testData{
		key: "key1",
		val: 4,
	}
	item2 := testData{
		key: "key2",
		val: 3,
	}
	entries := []Entry{
		{
			Item:   item1,
			Weight: 0.4,
		},
		{
			Item:   item2,
			Weight: 0.6,
		},
	}
	rand, err := NewWeightedRandom(entries)
	assert.Nil(t, err)
	retItem, err := rand.GetWithSeed("ab")
	assert.Nil(t, err)
	assert.Equal(t, item2, retItem)

	for i := 1; i <= 10; i++ {
		retItem, err := rand.GetWithSeed("hi")
		assert.Nil(t, err)
		assert.Equal(t, item1, retItem)
	}
}

func TestDeterministicWeightedFewZeroWeight(t *testing.T) {
	item1 := testData{
		key: "key1",
		val: 4,
	}
	item2 := testData{
		key: "key2",
		val: 3,
	}
	entries := []Entry{
		{
			Item:   item1,
			Weight: 0.4,
		},
		{
			Item: item2,
		},
	}
	rand, err := NewWeightedRandom(entries)
	assert.Nil(t, err)
	retItem, err := rand.GetWithSeed("ab")
	assert.Nil(t, err)
	assert.Equal(t, item1, retItem)

	for i := 1; i <= 10; i++ {
		retItem, err := rand.GetWithSeed("hi")
		assert.Nil(t, err)
		assert.Equal(t, item1, retItem)
	}
}

func TestDeterministicWeightedAllZeroWeights(t *testing.T) {
	item1 := testData{
		key: "sort_key1",
		val: 4,
	}
	item2 := testData{
		key: "sort_key2",
		val: 3,
	}
	entries := []Entry{
		{
			Item: item1,
		},
		{
			Item: item2,
		},
	}
	rand, err := NewWeightedRandom(entries)
	assert.Nil(t, err)
	retItem, err := rand.GetWithSeed("hi")
	assert.Nil(t, err)
	assert.Equal(t, item2, retItem)

	for i := 1; i <= 10; i++ {
		retItem, err := rand.GetWithSeed("ab")
		assert.Nil(t, err)
		assert.Equal(t, item1, retItem)
	}
}

func TestDeterministicWeightInvalidWeights(t *testing.T) {
	item1 := testData{
		key: "key1",
		val: 4,
	}
	item2 := testData{
		key: "key2",
		val: 3,
	}
	entries := []Entry{
		{
			Item:   item1,
			Weight: -3.0,
		},
		{
			Item: item2,
		},
	}
	_, err := NewWeightedRandom(entries)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "invalid weight -3.000000")
}
