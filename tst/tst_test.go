package tst

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNode(t *testing.T) {
	n := New()

	assert.False(t, n.Contains("foo"))

	n.Set("foo", nil)
	assert.True(t, n.Contains("foo"))
	assert.False(t, n.Contains("foob"))

	n.Set("foob", nil)
	assert.True(t, n.Contains("foo"))
	assert.True(t, n.Contains("foob"))
}
