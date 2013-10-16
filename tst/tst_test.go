package tst

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	T, F := true, false

	tree := New()

	assert.Equal(t, F, tree.Contains("foo"))

	tree.Set("foo", nil)
	{
		assert.Equal(t, F, tree.Contains("f"))
		assert.Equal(t, F, tree.Contains("fo"))
		assert.Equal(t, T, tree.Contains("foo"))
		assert.Equal(t, F, tree.Contains("foob"))

		assert.Nil(t, tree.Get("f"))
		assert.Nil(t, tree.Get("fo"))
		assert.Equal(t, "foo", tree.Get("foo"))
		assert.Nil(t, tree.Get("foob"))
	}

	tree.Set("fo", nil)
	{
		assert.Equal(t, F, tree.Contains("f"))
		assert.Equal(t, T, tree.Contains("fo"))
		assert.Equal(t, T, tree.Contains("foo"))
		assert.Equal(t, F, tree.Contains("foob"))

		assert.Nil(t, tree.Get("f"))
		assert.Equal(t, "fo", tree.Get("fo"))
		assert.Equal(t, "foo", tree.Get("foo"))
		assert.Nil(t, tree.Get("foob"))
	}

	tree.Set("foob", nil)
	{
		assert.Equal(t, F, tree.Contains("f"))
		assert.Equal(t, T, tree.Contains("fo"))
		assert.Equal(t, T, tree.Contains("foo"))
		assert.Equal(t, T, tree.Contains("foob"))

		assert.Nil(t, tree.Get("f"))
		assert.Equal(t, "fo", tree.Get("fo"))
		assert.Equal(t, "foo", tree.Get("foo"))
		assert.Equal(t, "foob", tree.Get("foob"))
	}

	tree.Set("f", nil)
	{
		assert.Equal(t, T, tree.Contains("f"))
		assert.Equal(t, T, tree.Contains("fo"))
		assert.Equal(t, T, tree.Contains("foo"))
		assert.Equal(t, T, tree.Contains("foob"))

		assert.Equal(t, "f", tree.Get("f"))
		assert.Equal(t, "fo", tree.Get("fo"))
		assert.Equal(t, "foo", tree.Get("foo"))
		assert.Equal(t, "foob", tree.Get("foob"))
	}

	//	vals := tree.AllWithPrefix("x", 0)
	//	assert.Equal(t, 0, len(vals), fmt.Sprintf("vals: %q", vals))

	//	vals = tree.AllWithPrefix("xoooooo", 0)
	//	assert.Equal(t, 0, len(vals), fmt.Sprintf("vals: %q", vals))

	//	vals = tree.AllWithPrefix("f", 0)
	//	assert.Equal(t, 4, len(vals), fmt.Sprintf("vals: %q", vals))

	//	vals = tree.AllWithPrefix("fo", 0)
	//	assert.Equal(t, 3, len(vals), fmt.Sprintf("vals: %q", vals))

	//	vals = tree.AllWithPrefix("foo", 0)
	//	assert.Equal(t, 2, len(vals), fmt.Sprintf("vals: %q", vals))

	//	vals = tree.AllWithPrefix("foob", 0)
	//	assert.Equal(t, 1, len(vals), fmt.Sprintf("vals: %q", vals))
}
