package kv

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeMemory(member string, init MemoryTable) *Client {
	return NewClient(member, NewMemoryDriver(init))
}

func TestKVCreateDelete(t *testing.T) {
	ctx := context.Background()
	c := makeMemory("basic", nil)

	c2 := makeMemory("notexist", nil)

	nonce, err := c.CreateUint64(ctx, "uint", 0)
	assert.Nil(t, err)

	_, err = c.CreateUint64(ctx, "uint", 0)
	assert.Equal(t, err, ErrAlreadySet)

	assert.Equal(t, c.Delete(ctx, "uint", []byte{1}), ErrNotEqual)
	assert.Nil(t, c.Delete(ctx, "uint", nonce))
	assert.Equal(t, c.Delete(ctx, "uint", nonce), ErrNotEqual)
	assert.Equal(t, c.Delete(ctx, "uint", nil), ErrUnsetValue)

	// XXX this tests another branch where the member table itself may not exist yet.
	assert.Equal(t, c2.Delete(ctx, "uint-notexist", nonce), ErrNotEqual)
	assert.Equal(t, c2.Delete(ctx, "uint-notexist", nil), ErrUnsetValue)

	_, err = c.CreateUint64(ctx, "uint", 0)
	assert.Nil(t, err)
	assert.Nil(t, c.Delete(ctx, "uint", nil))
}

func TestKVBasic(t *testing.T) {
	c := makeMemory("basic", nil)
	ctx := context.Background()
	basicTest(ctx, t, c)
}

func TestKVMember(t *testing.T) {
	ctx := context.Background()
	c1 := makeMemory("member1", nil)
	c2 := makeMemory("member2", nil)

	_, err := c1.SetUint64(ctx, "one", 1)
	assert.Nil(t, err)

	_, _, err = c2.GetUint64(ctx, "one")
	assert.Equal(t, err, ErrUnsetValue)
	_, err = c2.SetUint64(ctx, "one", 2)
	assert.Nil(t, err)

	out, _, err := c1.GetUint64(ctx, "one")
	assert.Nil(t, err)
	assert.Equal(t, uint64(1), out)
}
