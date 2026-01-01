# gml

Rendering can be slow due to extensive recursion and a large number of allocated objects.

One proposed improvement is to introduce a `gml.Root` type via a constructor:

```go
gml.New() gml.Root
```

The `gml.Root` instance would contain a global builder shared by all child elements, reducing allocations and improving performance.

Additionally, all HTML tags would become methods on `gml.Root` instead of global functions. For example:

```go
// Current global function
func A(_, _)

// Proposed method on gml.Root
func (r *gml.Root) A(_, _)
```

This approach centralizes the builder state and can lead to more efficient rendering.

