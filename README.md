# teivah/multilock

_multilock_ is a Go library allowing to store multiple `sync.Mutex` or `sync.RWMutex`.

The main benefit is to reduce the memory footprint if we need to have a mutex for every single structure of a set (e.g. thousand of maps or slices).

The internal data structure managed by _multilock_, depending on the use case, has either a fixed or a variable length.
Accessing a variable one means acquiring a shared lock first for every access (which does not exist for a fixed structure).

## Structures

* Fixed length structure of Mutex: `multilock.Fixed`
* Fixed length structure of RWMutex: `multilock.RWFixed`
* Variable length structure of Mutex: `multilock.Var`
* Variable length structure of RWMutex: `multilock.RWVar`

## Examples

```go
const multilockLength = 5
const customStructLength = 1000

// Initialize a fixed length multilock structure
mlock := multilock.NewFixed(multilockLength)

// Create custom structures
maps := make([]map[string]string, customStructLength)
for i := 0; i < 1000; i++ {
	maps[i] = make(map[string]string)
}

// Retrieve a lock for a given maps index
mutex := mlock.Get(maps[42])
mutex.Lock()
defer mutex.Unlock()
``` 