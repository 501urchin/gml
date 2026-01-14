package gmlx

import (
	"context"
	"sync"

	"github.com/501urchin/gml"
	"github.com/501urchin/gml/pkg"
	"github.com/cespare/xxhash/v2"
)

type store struct {
	mu    sync.RWMutex
	store map[uint64]string
}

var memoStore = store{store: make(map[uint64]string)}

// Memoize memoizes the rendered HTML for the lifetime of the process.
//
// WARNING: This function is UNBOUNDED.
// Each unique dependency value creates a new cache entry that is never
// evicted or expired. Memory usage will grow indefinitely as new
// dependency values are observed.
//
// This function must only be used for special cases where:
//   - The dependency space is known, small, and finite
//   - The render function is pure and deterministic
//   - The output is immutable for the lifetime of the process
//
// Do NOT use this for request-scoped data, user input, or any
// high-cardinality dependencies.
func Memoize[T any](ctx context.Context, dep T, elm func(dep T) gml.GmlElement) gml.GmlElement {
	cacheKey := xxhash.Sum64(pkg.StringToBytes(pkg.ConvertToString(dep)))

	memoStore.mu.RLock()
	v, hit := memoStore.store[cacheKey]
	memoStore.mu.RUnlock()

	if hit {
		return gml.Raw(v)
	}

	memoStore.mu.Lock()
	defer memoStore.mu.Unlock()

	html, err := elm(dep).RenderHtml(ctx)
	if err != nil {
		return elm(dep)
	}

	rhtml := string(html)

	memoStore.store[cacheKey] = rhtml

	return gml.Raw(rhtml)
}

// ClearMemo removes all memoized entries.
// This does not change the unbounded nature of Memoize and should only
// be used for explicit invalidation or testing.
func ClearMemo() {
	memoStore.mu.Lock()
	memoStore.store = make(map[uint64]string)
	memoStore.mu.Unlock()
}
