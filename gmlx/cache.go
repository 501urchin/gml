package gmlx

import (
	"context"
	"fmt"
	"sync"

	"github.com/501urchin/gml"
	"github.com/501urchin/gml/pkg"
	"github.com/cespare/xxhash/v2"
)

type store struct {
	mu    sync.RWMutex
	store map[uint64]string
}

var cacheStore = store{store: make(map[uint64]string)}

// this is only reccomended for componenets that have a deeply nested structure
func Cache[T any](ctx context.Context, dep T, elm func(dep T) gml.GmlElement) gml.GmlElement {
	cacheKey := xxhash.Sum64(pkg.StringToBytes(fmt.Sprintf("%v", dep))) // could optimize this further by implementing switch(type)

	cacheStore.mu.RLock()
	v, hit := cacheStore.store[cacheKey]
	cacheStore.mu.RUnlock()

	if hit {
		return gml.Raw(v)
	}

	html, err := elm(dep).RenderHtml(ctx)
	if err != nil {
		return elm(dep)
	}

	rhtml := string(html)

	cacheStore.mu.Lock()
	cacheStore.store[cacheKey] = rhtml
	cacheStore.mu.Unlock()

	return gml.Raw(rhtml)
}
