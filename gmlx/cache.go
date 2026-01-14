package gmlx

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"sync"

	"github.com/501urchin/gml"
	"github.com/cespare/xxhash/v2"
)

type store struct {
	mu    sync.RWMutex
	store map[uint64]string
}

var cacheStore = store{store: make(map[uint64]string)}

func binaryMarshalAnySlice(vals []any) []byte {
	buf := new(bytes.Buffer)

	for _, v := range vals {
		binary.Write(buf, binary.LittleEndian, v)
	}

	return buf.Bytes()
}

// this is only reccomended for componenets that have a deeply nested structure
func Cache(ctx context.Context, element gml.GmlElement, deps ...any) gml.GmlElement {
	jsonobj, err := json.Marshal(deps)
	if err != nil {
		return element
	}
	cacheKey := xxhash.Sum64(jsonobj)

	cacheStore.mu.RLock()
	v, hit := cacheStore.store[cacheKey]
	cacheStore.mu.RUnlock()

	if hit {
		return gml.Raw(v)
	}

	html, err := element.RenderHtml(ctx)
	if err != nil {
		return element
	}

	rhtml := string(html)

	cacheStore.mu.Lock()
	cacheStore.store[cacheKey] = rhtml
	cacheStore.mu.Unlock()

	return gml.Raw(rhtml)
}
