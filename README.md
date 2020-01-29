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

In the following example, we will create a 1/10 ratio multilock (10 mutexes to handle 100 maps):

```go
const multilockLength = 10
const customStructLength = 100

// Initialize a fixed length multilock structure
mlock := multilock.NewFixed(multilockLength)

// Create custom structures
maps := make([]map[string]string, customStructLength)
for i := 0; i < customStructLength; i++ {
	maps[i] = make(map[string]string)
}

// Retrieve a lock for a given maps index
mutex := mlock.Get(maps[42])
mutex.Lock()
defer mutex.Unlock()
``` 

Internally, _multilock_ has a distribution strategy which is basically hashing the key and returning a modulo based on the length provided.

It is also possible to override this distribution strategy this way:
```go
multilock.NewFixed(10, multilock.WithCustomDistribution(func(i interface{}, length int) int {
    // Return an int between 0 and 10 depending on our distribution strategy
}))
```