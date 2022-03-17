package atomic_test

import (
	"fmt"
	stdlibatomic "sync/atomic"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/alecthomas/atomic"
)

func TestValue(t *testing.T) {
	v := atomic.New("hello")
	assert.Equal(t, "hello", v.Load())
	v.Store("world")
	assert.Equal(t, "world", v.Load())
}

func TestValueZeroValue(t *testing.T) {
	var v atomic.Value[string]
	assert.Equal(t, "", v.Load())
	v.Store("world")
	assert.Equal(t, "world", v.Load())
}

func TestInt32(t *testing.T) {
	v := atomic.NewInt32(0)
	assert.Equal(t, 0, v.Load())
	assert.Equal(t, true, v.CompareAndSwap(0, 10))
	assert.Equal(t, false, v.CompareAndSwap(0, 10))
}

func BenchmarkInt64Add(b *testing.B) {
	v := atomic.NewInt64(0)
	for i := 0; i < b.N; i++ {
		v.Add(1)
	}
}

func BenchmarkIntInterfaceAdd(b *testing.B) {
	var v atomic.Int[int64] = atomic.NewInt64(0)
	for i := 0; i < b.N; i++ {
		v.Add(1)
	}
}

func BenchmarkStdlibInt64Add(b *testing.B) {
	var n int64
	for i := 0; i < b.N; i++ {
		stdlibatomic.AddInt64(&n, 1)
	}
}

func BenchmarkInterfaceStore(b *testing.B) {
	var v atomic.Interface[string] = atomic.New("hello")
	for i := 0; i < b.N; i++ {
		v.Store(fmt.Sprint(i))
	}
}

func BenchmarkValueStore(b *testing.B) {
	v := atomic.New("hello")
	for i := 0; i < b.N; i++ {
		v.Store(fmt.Sprint(i))
	}
}

func BenchmarkStdlibValueStore(b *testing.B) {
	v := stdlibatomic.Value{}
	for i := 0; i < b.N; i++ {
		v.Store(fmt.Sprint(i))
	}
}
